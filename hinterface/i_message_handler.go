package hinterface

type IMessageHandler interface {
	DoMessageHandler(IRequest)
	AddRouter(uint32, IRouter)
	StartWorkPoll()
	SendMsg2TaskQueue(IRequest)
}
