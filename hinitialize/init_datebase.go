package hinitialize

import (
	"github.com/jhinih/hin/hconfig"
	"github.com/jhinih/hin/hglobal"
	"github.com/jhinih/hin/hlog/zlog"
	"github.com/jhinih/hin/hpkg/database"
	"github.com/jhinih/hin/hpkg/mysqlx"
	"github.com/jhinih/hin/hpkg/redisx"
)

func InitDataBase(config hconfig.Config) {
	for _, name := range config.DB.Driver {
		switch name {
		case "mysql":
			database.InitDataBases(mysqlx.NewMySql(), config)
		default:
			//zlog.Fatalf("不支持的数据库驱动：%s", name)
		}
	}
}
func InitRedis(config hconfig.Config) {
	if config.Redis.Enable {
		var err error
		hglobal.Rdb, err = redisx.GetRedisClient(config)
		if err != nil {
			zlog.Errorf("无法初始化Redis : %v", err)
		}
		zlog.Infof("初始化Redis成功！")
	} else {
		zlog.Warnf("不使用Redis")
	}
}
