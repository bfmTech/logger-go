# logger-go

logger sdk go 版本。将日志内容按照统一格式通过控制台、文件、http上传到阿里云sls，支持缓存上传。


## 安装

```bash
go get -u github.com/bfmTech/logger-go
```

## 快速开始

可参考 examples 中的例子

**重点：Logger 申明为全局变量，初始化一次！可直接使用下面的logger.go**

logger/logger.go (根据情况替换[应用名称]和选择日志上传方式)
```go
package logger

import (
	"errors"
	"log"
	"sync"

	winner_logger "github.com/bfmTech/logger-go"
)

var logger winner_logger.Logger
var once sync.Once

func init() {
	initLogger()
}

func initLogger() winner_logger.Logger {
	var err error
	once.Do(func() {
		logger, err = winner_logger.NewLogger("应用名称", winner_logger.Console) // winner_logger.Console、winner_logger.File、winner_logger.Http
	})

	if err != nil {
		log.Fatal(errors.New("logger初始化失败：" + err.Error()))
	}

	return logger
}

/**
 * @description: 获取logger初始化
 * @return {*}
 */
func GetLogger() winner_logger.Logger {
	if logger == nil {
		return initLogger()
	}

	return logger
}
```

main.go (logger-demo 为 go.mod module)

```go
package main

import (
	"errors"
	"log"

	winner_logger "github.com/bfmTech/logger-go"
	"logger-demo/logger"
)

func main() {
    // 申明为全局变量
	logger := logger.GetLogger()
    // 程序最终退出时调用
	defer logger.Close()

	logger.Info("这是info消息1", "消息2", "消息3")
	logger.Error(errors.New("出错啦"))
	logger.Access(&winner_logger.AccessLog{})
}
```


## 详细说明

**日志存储时长：360天**

**日志上传方式**

|  上传方式   | 使用环境  | 说明  |
|  ----  | ----  | ----  |
| Console  | 容器平台(推荐) | 日志输出到控制台，Logtail采集 |
| File  | 容器平台 | 日志根据日期保存到指定目录，Logtail采集 |
| Http  | 任意环境(需配置阿里云相关环境变量) | http上传日志到阿里sls |


**日志类型说明**  
不同的方法的日志类型level不同 可在查询时筛选
```code
-logger.Debug    //支持多个参数，可根据每个参数检索
-logger.Info     //支持多个参数，可根据每个参数检索
-logger.Warn     //支持多个参数，可根据每个参数检索
-logger.Error    //只接受Error对象 （推送钉钉告警消息）
-logger.Access   //access日志，根据方法提示的参数传递
```


**access日志 字段说明**

|  字段   | 类型  | 说明  |
|  ----  | ----  | ----  |
| method  | string | 请求方法 |
| status  | number | HTTP请求状态 |
| beginTime  | number | 请求开始时间(秒时间戳) |
| endTime  | number | 请求结束时间(秒时间戳) |
| referer  | string | 请求来源 |
| httpHost  | string | 请求地址 |
| interface  | string | 请求接口 |
| reqQuery  | string | 请求url参数 |
| reqBody  | string | 请求body参数 |
| resBody  | string | 请求返回参数 |
| clientIp  | string | 客户端IP |
| userAgent  | string | 用户终端浏览器等信息  |
| reqId  | string | header[X-Request-ID]请求id，用于链路追踪 |
| headers  | string | 其他数据，比如：token |


**关于链路追踪的实现**  
* 获取请求header中的X-Request-ID (Nginx access日志中的req_id) 标记为req_id  
* 所有应用日志记录上req_id  
* 应用上下游调用时，透传 X-Request-ID  
* 可在监控系统中通过req_id在访问日志和应用日志查询到完成的请求链路  


**上传方式为Http时环境变量说明**
|  变量名   | 说明  |
|  ----  | ----  |
| LOGGER_ALIYUN_ENDPOINT  | 阿里云sls公网服务入口 |
| LOGGER_ALIYUN_PROJECTNAME  | 阿里云sls项目（Project） |
| LOGGER_ALIYUN_LOGSTORENAME  | 阿里云sls日志库（Logstore） |
| LOGGER_ALIYUN_ACCESSKEYID  | AccessKey ID，建议使用RAM用户的AccessKey信息。 |
| LOGGER_ALIYUN_ACCESSKEYSECRET  | AccessKey Secret，建议使用RAM用户的AccessKey信息。 |

## 注意
1、 `appName` 需唯一，且有意义，用于检索和报错时通知负责人。

2、日志上传方式为`File`或`Http`时，程序退出时必须调用Close()，否则可能导致最后部分日志丢失。

3、日志上传方式为`Http`时，需要配置阿里云相关环境变量，请联系管理员。

4、日志默认存储时长为**360**天，如有特殊需求请联系管理员。

5、日志上传方式为`File`时，可配置环境变量`NODE_APP_DATA`，设置文件存储目录。