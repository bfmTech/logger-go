package winner_logger

import (
	"fmt"
	"reflect"
	"strings"
)

type applicationLog struct {
	AppName string
	Level   string
	LogTime string
	Message string
	Stack   string
}

func (l *applicationLog) format() string {
	return fmt.Sprintf("[%s] [%s] [%s] [%s] %s", l.LogTime, l.Level, l.AppName, l.Stack, l.Message)
}

/**
 * @description: access日志结构体
 * @return {*}
 */
type AccessLog struct {
	Method    string
	Status    int32
	BeginTime int64
	EndTime   int64
	Referer   string
	HttpHost  string
	Interface string
	ReqQuery  string
	ReqBody   string
	ResBody   string
	ClientIp  string
	UserAgent string
	ReqId     string
	Headers   string
}

func (l *AccessLog) Format() string {
	var logStrArr []string
	v := reflect.ValueOf(*l)
	for i := 0; i < v.NumField(); i++ {
		logStrArr = append(logStrArr, fmt.Sprintf("%v", v.Field(i)))
	}

	return strings.Join(logStrArr, separator)
}
