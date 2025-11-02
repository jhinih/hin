package hinterface

type IServer interface {
	Start()
	Stop()
	Server()

	AddRouter(uint32, IRouter)
	GetConnectionManagerHandler() IConnectionManager

	SetConnectionStartHook(func(IConnection))
	SetConnectionStopHook(func(IConnection))

	CallConnectionStartHook(IConnection)
	CallConnectionStopHook(IConnection)
}
