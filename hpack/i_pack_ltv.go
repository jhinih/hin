package hpack

import "github.com/jhinih/hin/hinterface"

type ILTVPack interface {
	GetHeadLen() uint32
	Pack(message hinterface.IMessage) ([]byte, error)
	UnPack([]byte) (hinterface.IMessage, error)
}
