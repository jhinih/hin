package hinterface

type IServer interface {
	Start()
	Stop()
	Server()

	//// Get connection management (得到连接管理)
	//GetConnMgr() IConnManager
	//
	//// Routing feature: register a routing business method for the current service for client link processing use
	////(路由功能：给当前服务注册一个路由业务方法，供客户端连接处理使用)
	//AddRouter(msgID uint32, router IRouter)
	//
	//// Get the message processing module binding method for the Server
	//// (获取Server绑定的消息处理模块)
	//GetMsgHandler() IMsgHandle
}
