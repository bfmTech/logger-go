package file

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"unsafe"

	"github.com/bfmTech/logger-go/common"
)

type FileLogger struct {
	AppName         string
	maxBufferLength int64            // 最大缓存字符串长度
	maxBufferSize   int64            // 最大缓存日志条数
	bufferLog       []unsafe.Pointer // 缓存日志
	bufferLength    int64            // 缓存日志长度
	bufferChan      chan string
	t               *time.Ticker
	filePath        string // 文件保存路径
}

func (l *FileLogger) Initialize() error {
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
	l.maxBufferLength = 1 * 1024 * 1024
	l.maxBufferSize = 100
	l.bufferLog = make([]unsafe.Pointer, 0, 200)
	l.bufferLength = 0
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

func (l *FileLogger) createInterval() {
	for {
		select {
		case <-l.t.C:
			l.flush()
		case logStr := <-l.bufferChan:
			l.bufferLength += int64(len(logStr))
			l.bufferLog = append(l.bufferLog, unsafe.Pointer(&logStr))
			if len(l.bufferLog) >= int(l.maxBufferSize) || l.bufferLength >= l.maxBufferLength {
				l.flush()
			}
		}
	}
}

func (l *FileLogger) flush() error {
	if len(l.bufferLog) > 0 {
		tempBufferLog := l.bufferLog
		l.bufferLog = l.bufferLog[:0]
		l.bufferLength = 0

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

func (l *FileLogger) Debug(message ...string) {
	l.log(common.Debug, message)
}

func (l *FileLogger) Info(message ...string) {
	l.log(common.Info, message)
}

func (l *FileLogger) Warn(message ...string) {
	l.log(common.Warn, message)
}

func (l *FileLogger) Error(message error) {
	l.log(common.Error, []string{message.Error()})
}

func (l *FileLogger) Access(accessLog *common.AccessLog) {
	l.log(common.Access, []string{accessLog.Format()})
}

func (l *FileLogger) Close() {
	l.flush()
}

func (l *FileLogger) log(level common.Level, messages []string) {
	if len(messages) == 0 {
		return
	}

	logStr := common.GetApplicationLogStr(level, l.AppName, messages, 4)

	l.bufferChan <- logStr
}
