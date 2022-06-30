package common

type Level string

const (
	Debug  Level = "debug"
	Info   Level = "info"
	Warn   Level = "warn"
	Error  Level = "error"
	Access Level = "access"
)

type LoggerMethod string

const (
	Console LoggerMethod = "console"
	File    LoggerMethod = "file"
	Http    LoggerMethod = "http"
)

const Separator = " \t "
