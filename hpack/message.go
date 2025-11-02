package hpack

type Message struct {
	ID      uint32
	Data    []byte
	DataLen uint32
}

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		ID:      id,
		Data:    data,
		DataLen: uint32(len(data)),
	}
}

func (msg *Message) GetID() uint32 {
	return msg.ID
}
func (msg *Message) GetData() []byte {
	return msg.Data
}
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

func (msg *Message) SetID(id uint32) {
	msg.ID = id
}
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}
