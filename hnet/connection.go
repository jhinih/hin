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
	Name string

	Conn     net.Conn
	ConnID   uint32
	isClosed bool

	ExitChan chan bool

	read2WriteChan chan []byte

	MsgHandler hinterface.IMessageHandler

	property     map[string]interface{}
	propertyLock sync.RWMutex

	Pack hinterface.IPack

	ConnectionStartHook func(hinterface.IConnection)
	ConnectionStopHook  func(hinterface.IConnection)
	ConnectionManager   hinterface.IConnectionManager

	localAddr  string
	remoteAddr string
}

func NewServerConnection(server hinterface.IServer, conn net.Conn, connID uint32, MsgHandler hinterface.IMessageHandler) *Connection {
	c := &Connection{
		Name:           server.GetName(),
		Conn:           conn,
		ConnID:         connID,
		isClosed:       false,
		ExitChan:       make(chan bool, 1),
		read2WriteChan: make(chan []byte, 1),
		property:       make(map[string]interface{}),

		localAddr:  conn.LocalAddr().String(),
		remoteAddr: conn.RemoteAddr().String(),
	}

	c.Pack = server.GetPack()
	c.ConnectionManager = server.GetConnectionManager()
	c.MsgHandler = server.GetMsgHandler()
	c.ConnectionStartHook = server.GetConnectionStartHook()
	c.ConnectionStopHook = server.GetConnectionStopHook()

	c.ConnectionManager.Add(c)
	fmt.Println("[connection number " + strconv.Itoa(c.ConnectionManager.Len()) + "]")

	return c
}
func NewClientConnection(client hinterface.IClient, conn net.Conn) *Connection {
	c := &Connection{
		Name:           client.GetName(),
		Conn:           conn,
		isClosed:       false,
		ExitChan:       make(chan bool, 1),
		read2WriteChan: make(chan []byte, 1),
		property:       make(map[string]interface{}),

		localAddr:  conn.LocalAddr().String(),
		remoteAddr: conn.RemoteAddr().String(),
	}

	c.Pack = client.GetPack()
	c.MsgHandler = client.GetMsgHandler()
	c.ConnectionStartHook = client.GetConnectionStartHook()
	c.ConnectionStopHook = client.GetConnectionStartHook()

	return c
}
func (c *Connection) StartReader() {
	fmt.Println("[client", c.localAddr, "conn start reader]")
	defer fmt.Println("[client", c.RemoteAddress().String(), "conn reader exit]")
	defer c.Stop()

	for {
		select {
		case <-c.ExitChan:
			return
		default:
			head := make([]byte, c.Pack.GetHeadLen())
			_, err := io.ReadFull(c.Conn, head)
			if err != nil {
				fmt.Println("read head err:", err)
				break
			}

			msgHandler, err := c.Pack.UnPack(head)
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
	fmt.Println("[connection", strconv.Itoa(int(c.ConnID)), "start]")
	go c.StartReader()
	go c.StartWriter()
	c.CallOnConnectionStart()
}
func (c *Connection) Stop() {
	c.CallOnConnectionStop()
	c.ConnectionManager.Remove(c)
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
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn.(*net.TCPConn)
}
func (c *Connection) GetConnection() net.Conn {
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

	fmt.Println("[client", strconv.Itoa(int(c.ConnID)), "send", msgID, "data:", string(data))

	msg := hpack.NewMessage(msgID, data)
	binaryMsg, err := c.Pack.Pack(msg)
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
func (c *Connection) GetProperty(key string) interface{} {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()

	if value, ok := c.property[key]; ok {
		return value
	} else {
		return errors.New("property not exist")
	}
}
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}
func (c *Connection) CallOnConnectionStart() {
	if c.ConnectionStartHook != nil {
		fmt.Println("HIN CallOnConnectionStart....")
		c.ConnectionStartHook(c)
	}
}

func (c *Connection) CallOnConnectionStop() {
	if c.ConnectionStartHook != nil {
		fmt.Println("HIN CallOnConnectionStop....")
		c.ConnectionStopHook(c)
	}
}
