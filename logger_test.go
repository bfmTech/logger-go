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
		Method:        "get",
		Status:        200,
		TimeLocal:     "2022-06-23 09:32:12",
		RequestTime:   1,
		RequestUri:    "/fueling/main1?appid=jn8yixa2g9ladb1&appid=plgflredscsbrvy&merId=18&merId=18&t=9199316",
		Referer:       "app.epaynfc.com",
		RemoteAddr:    "47.98.74.125",
		BodyBytesSent: 21,
		UserAgent:     "Mozilla/5.0 (Linux; Android 9; COR-AL10 Build/HUAWEICOR-AL10; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/88.0.4324.93 Mobile Safari/537.36;psbc",
	})

	time.Sleep(time.Second * 3)
}
