package winner_logger

type logLevel string

const (
	debugLog  logLevel = "debug"
	infoLog   logLevel = "info"
	warnLog   logLevel = "warn"
	errorLog  logLevel = "error"
	accessLog logLevel = "access"
)

type loggerMethod string

const (
	Console loggerMethod = "console" // 控制台输出日志
	File    loggerMethod = "file"    // 文件记录日志
	Http    loggerMethod = "http"    // http同步日志
)

const separator = " \t "
