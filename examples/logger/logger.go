package logger

import (
	"errors"
	"log"
	"sync"

	winner_logger "github.com/bfmTech/logger-go"
)

var logger winner_logger.Logger
var once sync.Once

func init() {
	initLogger()
}

func initLogger() winner_logger.Logger {
	var err error
	once.Do(func() {
		logger, err = winner_logger.NewLogger("应用名称", winner_logger.Console) // winner_logger.Console、winner_logger.File、winner_logger.Http
	})

	if err != nil {
		log.Fatal(errors.New("logger初始化失败：" + err.Error()))
	}

	return logger
}

/**
 * @description: 获取logger初始化
 * @return {*}
 */
func GetLogger() winner_logger.Logger {
	if logger == nil {
		return initLogger()
	}

	return logger
}
