package hinitialize

import (
	"github.com/jhinih/hin/hglobal"
	"github.com/jhinih/hin/hlog/zlog"
	"runtime"
)

func Eve() {
	zlog.Warnf("开始释放资源！")
	errRedis := hglobal.Rdb.Close()
	if errRedis != nil {
		zlog.Errorf("Redis关闭失败 ：%v", errRedis.Error())
	}

	sqlDB, _ := hglobal.DB.DB()
	errDB := sqlDB.Close()
	if errDB != nil {
		zlog.Errorf("数据库关闭失败 ：%v", errDB.Error())
	}

	if hglobal.RabbitMQ != nil {
		if err := hglobal.RabbitMQ.Close(); err != nil {
			zlog.Errorf("RabbitMQ 关闭失败：%v", err)
		}
	}
	runtime.GC()
	if errDB == nil && errRedis == nil {
		zlog.Warnf("资源释放成功！")
	}
}
