package hinterface

type IPack interface {
	GetHeadLen() uint32
	Pack(message IMessage) ([]byte, error)
	UnPack([]byte) (IMessage, error)
}
