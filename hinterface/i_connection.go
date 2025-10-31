package hinterface

import (
	"context"
	"net"
)

// IConnection Define connection interface
type IConnection interface {
	// Start the connection, make the current connection start working
	// (启动连接，让当前连接开始工作)
	Start()
	// Stop the connection and end the current connection state
	// (停止连接，结束当前连接状态)
	Stop()

	// Returns ctx, used by user-defined go routines to obtain connection exit status
	// (返回ctx，用于用户自定义的go程获取连接退出状态)
	Context() context.Context

	GetName() string         // Get the current connection name (获取当前连接名称)
	GetConnection() net.Conn // Get the original socket from the current connection(从当前连接获取原始的socket)
	// Deprecated: use GetConnection instead
	GetTCPConnection() net.Conn // Get the original socket TCPConn from the current connection (从当前连接获取原始的socket TCPConn)
	GetConnID() uint64          // Get the current connection ID (获取当前连接ID)
	GetConnIdStr() string       // Get the current connection ID for string (获取当前字符串连接ID)
	GetMsgHandler() IMsgHandle  // Get the message handler (获取消息处理器)
	GetWorkerID() uint32        // Get Worker ID（获取workerID）

	// Send Message data directly to the remote TCP client (without buffering)
	// 直接将Message数据发送数据给远程的TCP客户端(无缓冲)
	SendMsg(msgID uint32, data []byte) error
}
