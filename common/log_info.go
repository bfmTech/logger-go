package common

import (
	"fmt"
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

type AccessLog struct {
	Method        string
	Status        int32
	TimeLocal     string
	RequestTime   float64
	RequestUri    string
	Referer       string
	RemoteAddr    string
	BodyBytesSent int64
	UserAgent     string
}

func (l *AccessLog) Format() string {
	return fmt.Sprintf("%s%s%v%s%v%s%s%s%s%s%s%s%v%s%s", l.Method, Separator, l.Status, Separator, l.RequestTime, Separator, l.RequestUri, Separator, l.Referer, Separator, l.RemoteAddr, Separator, l.BodyBytesSent, Separator, l.UserAgent)
}
