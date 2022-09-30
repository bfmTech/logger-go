package winner_logger

import (
	"fmt"
	"os"
)

type consoleLogger struct {
	AppName string
}

func (l *consoleLogger) initialize() error {
	return nil
}

func (l *consoleLogger) Debug(message ...string) {
	l.log(debugLog, message)
}

func (l *consoleLogger) Info(message ...string) {
	l.log(infoLog, message)
}

func (l *consoleLogger) Warn(message ...string) {
	l.log(warnLog, message)
}

func (l *consoleLogger) Error(message error) {
	l.log(errorLog, []string{message.Error()})
}

func (l *consoleLogger) Access(access *AccessLog) {
	l.log(accessLog, []string{access.Format()})
}

func (l *consoleLogger) Close() {
}

func (l *consoleLogger) SetStoringDays(days int) {
}

func (l *consoleLogger) log(level logLevel, messages []string) {
	if len(messages) == 0 {
		return
	}

	logStr := getApplicationLogStr(level, l.AppName, messages, 4)

	print(level, logStr)
}

func print(level logLevel, logStr string) {
	if level == errorLog {
		fmt.Fprintln(os.Stderr, logStr)
	} else {
		fmt.Println(logStr)
	}
}
