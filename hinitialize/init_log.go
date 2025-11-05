package hinitialize

import (
	"github.com/jhinih/hin/hconfig"
	"github.com/jhinih/hin/hlog"
	"github.com/jhinih/hin/hlog/zlog"
)

func InitLog(config *hconfig.Config) {
	logger := log.GetZap(config)
	zlog.InitLogger(logger)
}
