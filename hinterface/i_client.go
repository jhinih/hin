package hinterface

type IClient interface {
	Start()
	ReStart()
	Stop()
	AddRouter(uint32, IRouter)

	GetConnection() IConnection

	SetConnectionStartHook(func(IConnection))
	SetConnectionStopHook(func(IConnection))
	GetConnectionStartHook() func(IConnection)
	GetConnectionStopHook() func(IConnection)

	GetPack() IPack
	SetPack(IPack)
	GetMsgHandler() IMessageHandler

	SetName(string)
	GetName() string

	GetErrChan() <-chan error
}
