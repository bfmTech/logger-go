package logger

import (
	"github.com/bfmTech/logger-go"
	"github.com/bfmTech/logger-go/common"
)

/**
 * @description: logger初始化
 * @return {*}
 */
func InitLogger() (logger.Logger, error) {
	log, err := logger.NewLogger("logger-go-test", common.Console) // common.Console、common.File、common.Http
	if err != nil {
		return nil, err
	}

	return log, nil
}
