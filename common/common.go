package common

import (
	"runtime"
	"strconv"
	"strings"
	"time"
)

func GetStack(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	return file + ":" + strconv.Itoa(line)
}

func GetApplicationLogStr(level Level, appName string, messages []string, skipStack int) string {
	message := strings.Join(messages, Separator)

	now := time.Now().Format("2006-01-02 15:04:05.000")
	logInfo := &applicationLog{
		AppName: appName,
		Level:   string(level),
		LogTime: now,
		Stack:   GetStack(skipStack),
		Message: message,
	}

	return logInfo.format()
}
