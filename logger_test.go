package logger

import (
	"errors"
	"testing"
	"time"

	"github.com/bfmTech/logger-go/common"
)

func TestOne(t *testing.T) {
	log, _ := NewLogger("go_app", common.Console) // common.Console、common.File、common.Http
	defer log.Close()

	log.Info("这是info消息1", "这是info消息2")
	log.Debug("这是debug消息")
	log.Error(errors.New("这是error消息"))
	log.Access(&common.AccessLog{
		Method:    "get",
		Status:    200,
		BeginTime: 1657092964,
		EndTime:   1657092964,
		Referer:   "http://monitor.wngf.com.cn",
		HttpHost:  "monitor.wngf.com.cn",
		Interface: "/api/v2/warning/list",
		ReqQuery:  "page=1&limit=10",
		ReqBody:   "",
		ResBody:   "4e8eaca4d",
		ClientIp:  "113.132.211.1",
		UserAgent: "Mozilla/5.0 (Linux; Android 9; COR-AL10 Build/HUAWEICOR-AL10; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/88.0.4324.93 Mobile Safari/537.36;psbc",
		Headers:   "334kj3k4j3k4j",
	})

	time.Sleep(time.Second * 3)
}
