package hpack

import "github.com/jhinih/hin/hinterface"

type IPack interface {
	GetHeadLen() uint32
	Pack(message hinterface.IMessage) ([]byte, error)
	UnPack([]byte) (hinterface.IMessage, error)
}
