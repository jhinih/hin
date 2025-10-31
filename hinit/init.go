package initialize

import (
	"github.com/jhinih/hin/hglobal"
	"github.com/jhinih/hin/hutils"
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"syscall"
)

var (
	Ch *amqp.Channel
)

func Init() {
	// 初始化根目录
	InitPath()
	// 加载配置文件
	InitConfig()
	// 正式初始化日志
	InitLog(hglobal.Config)
	// 初始化数据库
	InitDataBase(*hglobal.Config)
	InitRedis(*hglobal.Config)

	// 初始化消息队列
	//InitMQ(*hglobal.Config)
	// 初始化全局雪花ID生成器
	InitSnowflake()
	// 开启定时任务
	Cron()
	// 初始化OSS服务
	InitOSS(*hglobal.Config)

	//// 初始化ElasticSearch
	//InitElasticsearch()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		hglobal.Stop()
	}()
}
func InitPath() {
	hglobal.Path = hutils.GetRootPath("")
}
