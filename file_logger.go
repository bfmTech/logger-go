package winner_logger

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"unsafe"
)

type fileLogger struct {
	AppName         string
	maxBufferLength int64            // 最大缓存日志条数
	maxBufferSize   int64            // 最大缓存字符串大小
	bufferLog       []unsafe.Pointer // 缓存日志
	bufferSize      int64            // 缓存日志字符串大小
	bufferChan      chan string
	t               *time.Ticker
	filePath        string // 文件保存路径
}

func (l *fileLogger) initialize() error {
	hostName, err := os.Hostname()
	if err != nil {
		return err
	}

	filePath := "/var/winnerlogs"

	/*
	 * 日志写入文件条件：
	 * 每隔3秒
	 * 文件字符长度大于等于1 * 1024 * 1024
	 * 日志条数大于等于100
	 */
	l.maxBufferSize = 1 * 1024 * 1024
	l.maxBufferLength = 100
	l.bufferLog = make([]unsafe.Pointer, 0, 200)
	l.bufferSize = 0
	l.bufferChan = make(chan string, 10000)
	l.t = time.NewTicker(time.Millisecond * 3000)
	l.filePath = fmt.Sprintf("%s/%s/%s/", filePath, l.AppName, hostName)

	err = os.MkdirAll(l.filePath, os.ModePerm)
	if err != nil {
		return err
	}

	go l.createInterval()

	return nil
}

func (l *fileLogger) createInterval() {
	for {
		select {
		case <-l.t.C:
			l.flush()
		case logStr := <-l.bufferChan:
			l.bufferSize += int64(len(logStr))
			l.bufferLog = append(l.bufferLog, unsafe.Pointer(&logStr))
			if len(l.bufferLog) >= int(l.maxBufferLength) || l.bufferSize >= l.maxBufferSize {
				l.flush()
			}
		}
	}
}

func (l *fileLogger) flush() error {
	if len(l.bufferLog) > 0 {
		tempBufferLog := l.bufferLog
		l.bufferLog = l.bufferLog[:0]
		l.bufferSize = 0

		// 失败重试3次
		retryNum := 2
		for {
			err := writeFile(l.filePath, &tempBufferLog)
			if err == nil {
				break
			}

			if retryNum == 0 {
				for _, v := range tempBufferLog {
					fmt.Println(*(*string)(v))
				}

				break
			}

			retryNum--
			time.Sleep(time.Second)
		}
	}

	return nil
}

func writeFile(filePath string, tempBufferLog *[]unsafe.Pointer) error {
	file, err := os.OpenFile(fmt.Sprintf("%s/logger-%s.log", filePath, time.Now().Format("2006-01-02")), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	for _, v := range *tempBufferLog {
		write.WriteString(*(*string)(v) + "\n")
	}
	err = write.Flush()

	return err
}

func (l *fileLogger) Debug(message ...string) {
	l.log(debugLog, message)
}

func (l *fileLogger) Info(message ...string) {
	l.log(infoLog, message)
}

func (l *fileLogger) Warn(message ...string) {
	l.log(warnLog, message)
}

func (l *fileLogger) Error(message error) {
	l.log(errorLog, []string{message.Error()})
}

func (l *fileLogger) Access(access *AccessLog) {
	l.log(accessLog, []string{access.Format()})
}

func (l *fileLogger) Close() {
	l.flush()
}

func (l *fileLogger) log(level logLevel, messages []string) {
	if len(messages) == 0 {
		return
	}

	logStr := getApplicationLogStr(level, l.AppName, messages, 4)

	l.bufferChan <- logStr
}
