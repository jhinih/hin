package hinitialize

import (
	"github.com/jhinih/hin/hconfig"
	"github.com/jhinih/hin/hlog/zlog"
	"github.com/jhinih/hin/hpkg/mq"
	"github.com/jhinih/hin/hpkg/mq/rabbitmqx"
)

func InitMQ(config hconfig.Config) {
	for _, name := range config.MQ.Enabled {
		switch name {
		case "rabbitmq":
			mq.InitMQ(rabbitmqx.NewRabbitMQ(), config)
		//case "kafka":
		//	mq.InitMQ(kafkax.NewKafKa(), config)
		default:
			zlog.Fatalf("不支持的消息队列驱动：%s", name)
		}
	}
}
