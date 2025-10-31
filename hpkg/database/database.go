package database

import (
	"github.com/jhinih/hin/hconfig"
	"github.com/jhinih/hin/hglobal"
	"github.com/jhinih/hin/hlog/zlog"
	"gorm.io/gorm"
)

type DataBase interface {
	GetDataBaseDsn(config hconfig.Config) string
	InitDataBase(config hconfig.Config) (*gorm.DB, error)
}

func InitDataBases(base DataBase, config hconfig.Config) {
	var err error
	hglobal.DB, err = base.InitDataBase(config)
	if err != nil {
		zlog.Fatalf("无法初始化数据库 %v", err)
		return
	}
	zlog.Infof("初始化数据库成功！")
	////对该数据库注册 hook
	//logic.RegisterHook(hglobal.DB)
	return
}
