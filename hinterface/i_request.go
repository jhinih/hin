package hinterface

type IRequest interface {
	GetConnection() IConnection
	GetMsgData() []byte
	GetMsgID() uint32
}
