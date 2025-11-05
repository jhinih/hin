package hinterface

import (
	"net"
)

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnection() net.Conn
	GetConnectionID() uint32
	// RemoteAddress 远程客户端IP
	RemoteAddress() net.Addr

	Send(uint32, []byte) error

	SetProperty(key string, value interface{})
	GetProperty(key string) interface{}
	RemoveProperty(key string)
}
