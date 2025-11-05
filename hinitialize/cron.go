package hinitialize

import (
	"github.com/jhinih/hin/hlog/zlog"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"time"
)

func Cron() {
	zone, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		logrus.Warn("加载时区错误:%v", err)
	}
	crontab := cron.New(cron.WithSeconds(), cron.WithLocation(zone))
	//// 每10分钟计算视频热度
	//_, err = crontab.AddFunc("@every 10m", ComputeVideoWeight)
	//if err != nil {
	//	zlog.Errorf("添加定时任务失败:%v", err)
	//}
	zlog.Infof("启动定时任务成功")
	crontab.Start()
}

//func ComputeVideoWeight() {
//	ctx := context.Background()
//	zlog.CtxInfof(ctx, "开始计算视频热度")
//	err := repository.NewVideoRequest(global.DB).ComputeVideoWeight()
//	if err != nil {
//		zlog.CtxErrorf(ctx, "计算视频热度失败:%v", err)
//	}
//	zlog.CtxInfof(ctx, "计算视频热度完成")
//}
