package hinterface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetTCpConnection() *net.TCPConn
	GetConnectionID() uint32
	// RemoteAddress 远程客户端IP
	RemoteAddress() net.Addr

	Send(uint32, []byte) error

	SetProperty(key string, value interface{})
	GetProperty(key string) (interface{}, error)
	RemoveProperty(key string)
}

type HandlerFunc func(*net.TCPConn, []byte, int) error
