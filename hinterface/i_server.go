package hinterface

type IServer interface {
	Start()
	Stop()
	Serve()
	GetName() string

	AddRouter(uint32, IRouter)

	SetConnectionStartHook(func(IConnection))
	SetConnectionStopHook(func(IConnection))
	GetConnectionStartHook() func(IConnection)
	GetConnectionStopHook() func(IConnection)

	GetPack() IPack
	SetPack(IPack)
	GetMsgHandler() IMessageHandler
	GetConnectionManager() IConnectionManager
}
