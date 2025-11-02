package hnet

import (
	"errors"
	"fmt"
	"github.com/jhinih/hin/hinterface"
	"github.com/jhinih/hin/hpack"
	"io"
	"net"
	"strconv"
	"sync"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnID   uint32
	isClosed bool

	ExitChan chan bool

	read2WriteChan chan []byte

	MsgHandler hinterface.IMessageHandler

	TcpServer hinterface.IServer

	property     map[string]interface{}
	propertyLock sync.RWMutex
}

func NewConnection(server hinterface.IServer, conn *net.TCPConn, connID uint32, MsgHandler hinterface.IMessageHandler) *Connection {
	return &Connection{
		Conn:           conn,
		ConnID:         connID,
		isClosed:       false,
		ExitChan:       make(chan bool, 1),
		read2WriteChan: make(chan []byte, 1),
		MsgHandler:     MsgHandler,
		TcpServer:      server,
		property:       make(map[string]interface{}),
	}
}
func (c *Connection) StartReader() {
	fmt.Println("[client", c.RemoteAddress().String(), "conn start reader]")
	defer fmt.Println("[client", c.RemoteAddress().String(), "conn reader exit]")
	defer c.Stop()

	for {
		p := hpack.NewTLVPack()

		head := make([]byte, p.GetHeadLen())
		_, err := io.ReadFull(c.Conn, head)
		if err != nil {
			fmt.Println("read head err:", err)
			break
		}

		msgHandler, err := p.UnPack(head)
		if err != nil {
			fmt.Println("unpack err:", err)
			break
		}

		if msgHandler.GetDataLen() > 0 {
			data := make([]byte, msgHandler.GetDataLen())
			if _, err := io.ReadFull(c.Conn, data); err != nil {
				fmt.Println("read data err:", err)
				break
			}
			msgHandler.SetData(data)
		}
		request := NewRequest(c, msgHandler)

		if 1 == 1 /*是否工作池*/ {
			go c.MsgHandler.SendMsg2TaskQueue(request)
		} else {
			go c.MsgHandler.DoMessageHandler(request)
		}

	}

}
func (c *Connection) StartWriter() {
	fmt.Println("[client", c.RemoteAddress().String(), "conn start writer]")
	defer fmt.Println("[client", c.RemoteAddress().String(), "conn writer exit]")
	//defer c.Stop()

	for {
		select {
		case data := <-c.read2WriteChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("writer err:", err)
				return
			}
		case <-c.ExitChan:
			return
		}
	}
}

func (c *Connection) Start() {
	c.TcpServer.GetConnectionManagerHandler().Add(c)
	fmt.Println("[connection number " + strconv.Itoa(c.TcpServer.GetConnectionManagerHandler().Len()) + "]")
	fmt.Println("[connection", strconv.Itoa(int(c.ConnID)), "start]")
	go c.StartReader()
	go c.StartWriter()
	c.TcpServer.CallConnectionStartHook(c)
}
func (c *Connection) Stop() {
	c.TcpServer.CallConnectionStopHook(c)
	c.TcpServer.GetConnectionManagerHandler().Remove(c)
	fmt.Println("[connection", strconv.Itoa(int(c.ConnID)), "stop]")
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	c.Conn.Close()
	c.ExitChan <- true

	close(c.ExitChan)
	close(c.read2WriteChan)
}
func (c *Connection) GetTCpConnection() *net.TCPConn {
	return c.Conn
}
func (c *Connection) GetConnectionID() uint32 {
	return c.ConnID
}
func (c *Connection) RemoteAddress() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(msgID uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("connection " + strconv.Itoa(int(c.ConnID)) + " is closed")
	}

	p := hpack.NewTLVPack()
	msg := hpack.NewMessage(msgID, data)
	binaryMsg, err := p.Pack(msg)
	if err != nil {
		return errors.New("pack err")
	}

	c.read2WriteChan <- binaryMsg
	return nil
}

func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	c.property[key] = value

}
func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()

	if value, ok := c.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("property not found")
	}
}
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}
