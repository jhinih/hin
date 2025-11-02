package hinterface

import (
	"net/http"
	"net/url"
)

type IClient interface {
	Restart()
	Start()
	Stop()
	AddRouter(msgID uint32, router IRouter)
	Conn() IConnection

	// SetOnConnStart Set the Hook function to be called when a connection is created for this Client
	// (设置该Client的连接创建时Hook函数)
	SetOnConnStart(func(IConnection))

	// SetOnConnStop Set the Hook function to be called when a connection is closed for this Client
	// (设置该Client的连接断开时的Hook函数)
	SetOnConnStop(func(IConnection))

	// GetOnConnStart Get the Hook function that is called when a connection is created for this Client
	// (获取该Client的连接创建时Hook函数)
	GetOnConnStart() func(IConnection)

	// GetOnConnStop Get the Hook function that is called when a connection is closed for this Client
	// (设置该Client的连接断开时的Hook函数)
	GetOnConnStop() func(IConnection)

	// GetMsgHandler Get the message handling module bound to this Client
	// (获取Client绑定的消息处理模块)
	GetMsgHandler() IMessageHandler

	// Get the error channel for this Client 获取客户端错误管道
	GetErrChan() <-chan error

	// Set the name of this Clien
	// 设置客户端Client名称
	SetName(string)

	// Get the name of this Client
	// 获取客户端Client名称
	GetName() string

	SetUrl(url *url.URL)

	GetUrl() *url.URL

	// Set custom headers for WebSocket connection
	// 设置WebSocket连接的自定义请求头
	SetWsHeader(http.Header)

	// Get custom headers for WebSocket connection
	// 获取WebSocket连接的自定义请求头
	GetWsHeader() http.Header
}
