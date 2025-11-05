package hglobal

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/jhinih/hin/hconfig"
	"github.com/jhinih/hin/hpkg/mq"
	"github.com/jhinih/hin/hutils/snowflakeUtils"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Path   string
	DB     *gorm.DB
	Rdb    *redis.Client
	Config *hconfig.Config
	//MQ       *amqp.Connection
	RabbitMQ mq.MQ
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
