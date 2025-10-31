package initialize

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jhinih/hin/hconfig"
	"github.com/jhinih/hin/hglobal"
	"github.com/jhinih/hin/hlog/zlog"
)

func InitOSS(config hconfig.Config) {
	// 初始化OSS客户端
	client, err := oss.New(config.Oss.Endpoint, config.Oss.AccessKeyID, config.Oss.AccessKeySecret)
	if err != nil {
		//zlog.Errorf("oss初始化失败 %v", err)
		panic(err)
	}
	// 获取Bucket
	bucket, err := client.Bucket(config.Oss.BucketName)
	if err != nil {
		zlog.Errorf("oss初始化失败 %v", err)
		panic(err)
	}
	hglobal.OssClient = client
	hglobal.OssBucket = bucket
	return
}
