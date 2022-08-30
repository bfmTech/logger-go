package http

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/bfmTech/logger-go/common"

	"github.com/aliyun/aliyun-log-go-sdk/producer"
)

type HttpLogger struct {
	AppName          string
	endpoint         string
	projectName      string
	logStoreName     string
	accessKeyId      string
	accessKeySecret  string
	producerInstance *producer.Producer
}

func (l *HttpLogger) Initialize() error {
	// 加载阿里云日志服务器相关配置的环境变量
	endpoint := os.Getenv("LOGGER_ALIYUN_ENDPOINT")
	if endpoint == "" {
		endpoint = "cn-hangzhou.log.aliyuncs.com"
	}
	projectName := os.Getenv("LOGGER_ALIYUN_PROJECTNAME")
	if projectName == "" {
		projectName = "k8s-log-custom-zwdfroh2"
	}
	logStoreName := os.Getenv("LOGGER_ALIYUN_LOGSTORENAME")
	if logStoreName == "" {
		logStoreName = "config-operation-log"
	}
	accessKeyId := os.Getenv("LOGGER_ALIYUN_ACCESSKEYID")
	if accessKeyId == "" {
		return errors.New("invalid env LOGGER_ALIYUN_ACCESSKEYID")
	}
	accessKeySecret := os.Getenv("LOGGER_ALIYUN_ACCESSKEYSECRET")
	if accessKeySecret == "" {
		return errors.New("invalid env LOGGER_ALIYUN_ACCESSKEYSECRET")
	}

	l.endpoint = endpoint
	l.projectName = projectName
	l.logStoreName = logStoreName
	l.accessKeyId = accessKeyId
	l.accessKeySecret = accessKeySecret

	// 启动producer实例
	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = l.endpoint
	producerConfig.AccessKeyID = l.accessKeyId
	producerConfig.AccessKeySecret = l.accessKeySecret
	producerConfig.AllowLogLevel = "error"
	l.producerInstance = producer.InitProducer(producerConfig)
	l.producerInstance.Start()

	return nil
}

func (l *HttpLogger) Debug(message ...string) {
	l.log(common.Debug, message)
}

func (l *HttpLogger) Info(message ...string) {
	l.log(common.Info, message)
}

func (l *HttpLogger) Warn(message ...string) {
	l.log(common.Warn, message)
}

func (l *HttpLogger) Error(message error) {
	l.log(common.Error, []string{message.Error()})
}

func (l *HttpLogger) Access(accessLog *common.AccessLog) {
	l.log(common.Access, []string{accessLog.Format()})
}

func (l *HttpLogger) Close() {
	l.producerInstance.SafeClose()
}

func (l *HttpLogger) log(level common.Level, messages []string) {
	if len(messages) == 0 {
		return
	}

	logStr := common.GetApplicationLogStr(level, l.AppName, messages, 4)

	log := producer.GenerateLog(uint32(time.Now().Unix()), map[string]string{"content": logStr})

	// 失败重试3次
	retryNum := 2
	for {
		err := l.producerInstance.SendLogWithCallBack(l.projectName, l.logStoreName, l.AppName, "", log, &Callback{logStr: logStr})
		if err == nil {
			break
		}

		if retryNum == 0 {
			fmt.Println(logStr)

			break
		}

		retryNum--
		time.Sleep(time.Second)
	}
}

/**
 * 发送日志的回调
 */
type Callback struct {
	logStr string
}

func (callback *Callback) Success(result *producer.Result) {
}

func (callback *Callback) Fail(result *producer.Result) {
	if !result.IsSuccessful() {
		fmt.Println(callback.logStr)
	}
}
