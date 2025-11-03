package api

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"github.com/golang/protobuf/proto"
	"github.com/jhinih/hin/demo/mmo_game/core"
	"github.com/jhinih/hin/demo/mmo_game/pb"
	"github.com/jhinih/hin/demo/mmo_game/utils/aiutils/jhinih_model/chatmodel"
	"github.com/jhinih/hin/hinterface"
	"github.com/jhinih/hin/hnet"
	"log"
)

// WorldChatApi World chat route business
// (WorldChatApi 世界聊天 路由业务)
type WorldChatApi struct {
	hnet.BaseRouter
}

func (*WorldChatApi) Handle(request hinterface.IRequest) {
	// 1. Decode the incoming proto protocol from the client
	// (1. 将客户端传来的proto协议解码)
	msg := &pb.Talk{}
	err := proto.Unmarshal(request.GetMsgData(), msg)
	if err != nil {
		fmt.Println("Talk Unmarshal error ", err)
		return
	}

	// 2. Identify which player sent the current message, retrieve from the connection property pID
	// (2. 得知当前的消息是从哪个玩家传递来的,从连接属性pID中获取)
	pID, err := request.GetConnection().GetProperty("pID")

	if err != nil {
		fmt.Println("GetProperty pID error", err)
		request.GetConnection().Stop()
		return
	}
	// 3. Get the player object based on pID
	// (3. 根据pID得到player对象)
	player := core.WorldMgrObj.GetPlayerByPID(pID.(int32))

	if pID.(int32)%2 == 0 {
		go func() {
			ctx := context.Background()
			model := chatmodel.NewArkChatModel(ctx)
			template := prompt.FromMessages(
				schema.FString,
				schema.SystemMessage("你是一个{role}"),
				&schema.Message{
					Role:    schema.User,
					Content: "{task}",
				},
			)
			params := map[string]any{
				"role": "猫娘",
				"task": msg.Content,
			}
			massage, err := template.Format(ctx, params)
			if err != nil {
				log.Fatal(err)
			}
			response, err := model.Generate(ctx, massage)
			if err != nil {
				log.Fatal(err)
			}

			AIplayer := core.WorldMgrObj.GetPlayerByPID(pID.(int32) - 1)
			AIplayer.Talk(response.Content)
		}()
	}
	// 4. Have the player object initiate the chat broadcast request
	// (4. 让player对象发起聊天广播请求)

	player.Talk(msg.Content)

}
