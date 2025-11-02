package hinterface

type IMessage interface {
	GetID() uint32
	GetData() []byte
	GetDataLen() uint32

	SetID(id uint32)
	SetData(data []byte)
	SetDataLen(len uint32)
}
