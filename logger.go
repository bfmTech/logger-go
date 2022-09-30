package winner_logger

/**
 * @description: 日志
 * @return {*}
 */
type Logger interface {
	initialize() error
	/**
	* @description: 记录debug日志
	* @param {...string} message 日志内容
	* @return {*}
	 */
	Debug(message ...string)
	/**
	* @description: 记录info日志
	* @param {...string} message 日志内容
	* @return {*}
	 */
	Info(message ...string)
	/**
	* @description: 记录warn日志
	* @param {...string} message 日志内容
	* @return {*}
	 */
	Warn(message ...string)
	/**
	* @description: 记录error日志
	* @param {error} message error对象
	* @return {*}
	 */
	Error(message error)
	/**
	* @description: 记录access日志
	* @param {*AccessLog} accessLog accessLog对象
	* @return {*}
	 */
	Access(accessLog *AccessLog)
	/**
	* @description: 上传当前缓存日志，释放资源
	* @return {*}
	 */
	Close()
	/**
	* @description: 设置文件日志存储天数
	* @return {*}
	 */
	SetStoringDays(days int)
}

/**
 * @description: 创建Logger对象
 * @param {string} appName 应用名
 * @param {common.LoggerMethod} method 日志记录方式 console、file、http
 * @return {*}
 */
func NewLogger(appName string, method loggerMethod) (Logger, error) {
	var log Logger

	switch method {
	case Console:
		log = &consoleLogger{AppName: appName}
	case File:
		log = &fileLogger{AppName: appName}
	case Http:
		log = &httpLogger{AppName: appName}
	default:
		log = &consoleLogger{AppName: appName}
	}

	err := log.initialize()

	return log, err
}
