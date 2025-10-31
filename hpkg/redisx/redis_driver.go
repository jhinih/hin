package redisx

import (
	"context"
	"fmt"
	"github.com/jhinih/hin/hconfig"
	"github.com/jhinih/hin/hlog/zlog"
	"github.com/redis/go-redis/v9"
)

const (
	redisAddr = "%s:%d"
)

func GetRedisClient(config hconfig.Config) (*redis.Client, error) {
	if !config.Redis.Enable {
		zlog.Warnf("不使用Redis模式")
		return nil, nil
	}
	client := redis.NewClient(&redis.Options{
		Network:         "",
		Addr:            fmt.Sprintf(redisAddr, hconfig.Conf.Redis.Host, hconfig.Conf.Redis.Port),
		Dialer:          nil,
		OnConnect:       nil,
		Username:        "",
		Password:        hconfig.Conf.Redis.Password,
		DB:              hconfig.Conf.Redis.DB,
		MaxRetries:      0,
		MinRetryBackoff: 0,
		MaxRetryBackoff: 0,
		DialTimeout:     0,
		ReadTimeout:     0,
		WriteTimeout:    0,
		PoolFIFO:        false,
		PoolSize:        1000,
		MinIdleConns:    1,
		PoolTimeout:     0,
		TLSConfig:       nil,
		Limiter:         nil,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		zlog.Fatalf("redis无法链接 %v", err)
		return nil, err
	}
	return client, nil
}
