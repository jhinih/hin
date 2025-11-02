package hnet

import "github.com/jhinih/hin/hinterface"

type Request struct {
	Connection hinterface.IConnection
	Msg        hinterface.IMessage
}

func NewRequest(connection hinterface.IConnection, msg hinterface.IMessage) *Request {
	return &Request{
		Connection: connection,
		Msg:        msg,
	}
}

func (r *Request) GetConnection() hinterface.IConnection {
	return r.Connection
}
func (r *Request) GetMsgData() []byte {
	return r.Msg.GetData()
}
func (r *Request) GetMsgID() uint32 {
	return r.Msg.GetID()
}
