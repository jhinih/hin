package global

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/jhinih/hin/demo/mmo_game/configs"
	"github.com/jhinih/hin/demo/mmo_game/utils/snowflakeUtils"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Path   string
	DB     *gorm.DB
	Rdb    *redis.Client
	Config *configs.Config
	//MQ       *amqp.Connection
	//RabbitMQ mq.MQ
	ESClient *elasticsearch.Client

	SnowflakeNode *snowflakeUtils.Node // 默认雪花ID生成节点
)

var (
	OssClient *oss.Client
	OssBucket *oss.Bucket
)
var ctxStop, cancelStop = context.WithCancel(context.Background())

func CtxDone() <-chan struct{} { return ctxStop.Done() }
func Stop()                    { cancelStop() }
