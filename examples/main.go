package main

import (
	"errors"
	"log"

	"github.com/bfmTech/logger-go/common"
	"github.com/bfmTech/logger-go/examples/logger"
)

func main() {
	logger, err := logger.InitLogger()
	if err != nil {
		log.Fatal("logger初始化失败：" + err.Error())
	}
	defer logger.Close()

	logger.Info("这是info消息1", "消息2", "消息3")

	logger.Error(errors.New("出错啦"))

	logger.Access(&common.AccessLog{
		Method:    "get",
		Status:    200,
		BeginTime: 1657092964,
		EndTime:   1657092964,
		Referer:   "http://xxx.test.com.cn",
		HttpHost:  "xx.test.com.cn",
		Interface: "/api/v2/warning/list",
		ReqQuery:  "page=1&limit=10",
		ReqBody:   "",
		ResBody:   "{\"code\":0,\"data\": \"\",\"msg\":\"success\"}",
		ClientIp:  "113.132.211.1",
		UserAgent: "Mozilla/5.0 (Linux; Android 9; COR-AL10 Build/HUAWEICOR-AL10; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/88.0.4324.93 Mobile Safari/537.36;psbc",
		ReqId:     "j4k34423kl3k4f5lk234js9",
		Headers:   "token:439skf2dk234",
	})

}
