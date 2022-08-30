package logger

import (
	winner_logger "github.com/bfmTech/logger-go"
)

/**
 * @description: logger初始化
 * @return {*}
 */
func InitLogger() (winner_logger.Logger, error) {
	log, err := winner_logger.NewLogger("logger-go-test", winner_logger.Console) // common.Console、common.File、common.Http
	if err != nil {
		return nil, err
	}

	return log, nil
}
