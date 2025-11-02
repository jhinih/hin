package hpack

import (
	"bytes"
	"encoding/binary"
	"github.com/jhinih/hin/hinterface"
)

type TLVPack struct {
}

func NewTLVPack() *TLVPack {
	return &TLVPack{}
}
func (p *TLVPack) GetHeadLen() uint32 {
	//ID(uint32-->4)+DataLen(uint32-->4)==8
	return 8
}

func (p *TLVPack) Pack(message hinterface.IMessage) ([]byte, error) {
	databuff := bytes.NewBuffer([]byte{})
	if err := binary.Write(databuff, binary.LittleEndian, message.GetID()); err != nil {
		return nil, err
	}
	if err := binary.Write(databuff, binary.LittleEndian, message.GetDataLen()); err != nil {
		return nil, err
	}

	if err := binary.Write(databuff, binary.LittleEndian, message.GetData()); err != nil {
		return nil, err
	}

	return databuff.Bytes(), nil
}
func (p *TLVPack) UnPack(binaryData []byte) (hinterface.IMessage, error) {
	databuff := bytes.NewBuffer(binaryData)
	msg := &Message{}
	if err := binary.Read(databuff, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}
	if err := binary.Read(databuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(databuff, binary.LittleEndian, &msg.Data); err != nil {
		return nil, err
	}

	return msg, nil
}
