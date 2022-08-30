# logger-go

logger sdk go 版本。将日志内容按照统一格式通过控制台、文件、http上传到阿里云sls，支持缓存上传。


## 安装

```bash
go get -u github.com/bfmTech/logger-go
```

## 快速开始

可参考 examples 中的例子

**重点：Logger 申明为全局变量，初始化一次！！！**

logger/logger.go (根据情况替换 [应用名称logger-go-test]和选择日志上传方式)
```go
package logger

import (
	"github.com/bfmTech/logger-go"
	"github.com/bfmTech/logger-go/common"
)

/**
 * @description: logger初始化
 * @return {*}
 */
func InitLogger() (logger.Logger, error) {
	log, err := logger.NewLogger("logger-go-test", common.Console) // common.Console、common.File、common.Http
	if err != nil {
		return nil, err
	}

	return log, nil
}
```

main.go (logger-demo 为 go.mod module)

```go
package main

import (
	"errors"
	"log"

	"github.com/bfmTech/logger-go/common"
	"logger-demo/logger"
)

func main() {
    // 申明为全局变量
	logger, err := logger.InitLogger()
	if err != nil {
		log.Fatal("logger初始化失败：" + err.Error())
	}
    // 程序最终退出时调用
	defer logger.Close()

	logger.Info("这是info消息1", "消息2", "消息3")
	logger.Error(errors.New("出错啦"))
	logger.Access(&common.AccessLog{})
}
```


## 详细说明

日志上传方式

|  LogType   | 使用环境  | 说明  |
|  ----  | ----  | ----  |
| common.Console  | 容器平台(推荐) | 日志输出到控制台，Logtail采集 |
| common.File  | 容器平台 | 日志根据日期保存到指定目录，Logtail采集 |
| common.Http  | 任意环境(需配置阿里云相关环境变量) | http上传日志到阿里sls |


日志类型 level 可在查询时筛选
```code
-level.debug    //支持多个参数，可根据每个参数检索
-level.info     //支持多个参数，可根据每个参数检索
-level.warn     //支持多个参数，可根据每个参数检索
-level.error    //只接受Error对象 （推送钉钉告警消息）
-level.access   //access日志，根据方法提示的参数传递
```

access日志 字段说明

|  字段   | 类型  | 说明  |
|  ----  | ----  | ----  |
| method  | string | 请求方法 |
| status  | number | HTTP请求状态 |
| beginTime  | number | 请求开始时间(秒时间戳) |
| endTime  | number | 请求结束时间(秒时间戳) |
| referer  | string | 请求来源 |
| httpHost  | string | 请求地址 |
| _interface  | string | 请求接口 |
| reqQuery  | string | 请求url参数 |
| reqBody  | string | 请求body参数 |
| resBody  | string | 请求返回参数 |
| clientIp  | string | 客户端IP |
| userAgent  | string | 用户终端浏览器等信息  |
| reqId  | string | header[X-Request-ID]请求id，用于链路跟踪 |
| headers  | string | 其他数据，比如：token |

## 注意
1、 `appName` 需唯一，且有意义，用于检索和报错时通知负责人。

2、日志上传方式为`File`或`Http`时，程序退出时必须调用close()，否则可能导致最后部分日志丢失。

3、日志上传方式为`Http`时，需要配置阿里云相关环境变量，请联系架构部。
