package winner_logger

import (
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

func getStack(level logLevel, skip int) string {
	var strStack string
	if level == errorLog {
		stack := debug.Stack()
		strStack = string(stack)
	} else {
		_, file, line, _ := runtime.Caller(skip)
		strStack = file + ":" + strconv.Itoa(line)
	}

	strStack = strings.ReplaceAll(strings.ReplaceAll(strStack, "[", ""), "]", "")
	return strStack
}

func getApplicationLogStr(level logLevel, appName string, messages []string, skipStack int) string {
	message := strings.Join(messages, separator)

	now := time.Now().Format("2006-01-02 15:04:05.000")
	logInfo := &applicationLog{
		AppName: appName,
		Level:   string(level),
		LogTime: now,
		Stack:   getStack(level, skipStack),
		Message: message,
	}

	return logInfo.format()
}
