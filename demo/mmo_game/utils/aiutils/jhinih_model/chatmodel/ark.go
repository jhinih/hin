package chatmodel

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/jhinih/hin/demo/mmo_game/log/zlog"
)

func NewArkChatModel(ctx context.Context) *ark.ChatModel {
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: "ab6875b4-cca1-404c-a371-70a7683929b7",
		Model:  "doubao-1-5-pro-32k-250115",
	})
	if err != nil {
		zlog.CtxErrorf(ctx, "创建ark模型失败%v", err)
	}
	return model
}
