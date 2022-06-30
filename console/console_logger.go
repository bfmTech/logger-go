package console

import (
	"fmt"
	"os"

	"github.com/bfmTech/logger-go/common"
)

type ConsoleLogger struct {
	AppName string
}

func (l *ConsoleLogger) Initialize() error {
	return nil
}

func (l *ConsoleLogger) Debug(message ...string) {
	l.log(common.Debug, message)
}

func (l *ConsoleLogger) Info(message ...string) {
	l.log(common.Info, message)
}

func (l *ConsoleLogger) Warn(message ...string) {
	l.log(common.Warn, message)
}

func (l *ConsoleLogger) Error(message error) {
	l.log(common.Error, []string{message.Error()})
}

func (l *ConsoleLogger) Access(accessLog *common.AccessLog) {
	l.log(common.Access, []string{accessLog.Format()})
}

func (l *ConsoleLogger) Close() {
}

func (l *ConsoleLogger) log(level common.Level, messages []string) {
	if len(messages) == 0 {
		return
	}

	logStr := common.GetApplicationLogStr(level, l.AppName, messages, 4)

	print(level, logStr)
}

func print(level common.Level, logStr string) {
	if level == common.Error {
		fmt.Fprintln(os.Stderr, logStr)
	} else {
		fmt.Println(logStr)
	}
}
