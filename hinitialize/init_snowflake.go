package hinitialize

import (
	"github.com/jhinih/hin/hglobal"
	"github.com/jhinih/hin/hlog/zlog"
	"github.com/jhinih/hin/hutils/snowflakeUtils"
)

func InitSnowflake() {
	var err error
	hglobal.SnowflakeNode, err = snowflakeUtils.NewNode(hglobal.DEFAULT_NODE_ID)
	if err != nil {
		zlog.Errorf("初始化雪花ID生成节点失败: %v", err)
		panic(err)
	}
}
