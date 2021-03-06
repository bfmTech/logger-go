package logger

import (
	"github.com/bfmTech/logger-go/common"
	"github.com/bfmTech/logger-go/console"
	"github.com/bfmTech/logger-go/file"
	"github.com/bfmTech/logger-go/http"
)

type Logger interface {
	Initialize() error
	Debug(message ...string)
	Info(message ...string)
	Warn(message ...string)
	Error(message error)
	Access(accessLog *common.AccessLog)
	Close()
}

func NewLogger(appName string, method common.LoggerMethod) (Logger, error) {
	var log Logger

	switch method {
	case common.Console:
		log = &console.ConsoleLogger{AppName: appName}
	case common.File:
		log = &file.FileLogger{AppName: appName}
	case common.Http:
		log = &http.HttpLogger{AppName: appName}
	default:
		log = &console.ConsoleLogger{AppName: appName}
	}

	err := log.Initialize()

	return log, err
}
