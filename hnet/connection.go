package hnet

import (
	"context"
	"encoding/hex"
	"errors"
	"github.com/jhinih/hin/hconfig"
	"github.com/jhinih/hin/hglobal"
	"github.com/jhinih/hin/hinterface"
	"github.com/jhinih/hin/hlog/zlog"
	"golang.org/x/net/websocket"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// CallBackFunc defines the callback function type
// (定义回调函数类型)
type CallBackFunc func()

// Connection TCP connection module
// Used to handle the read and write business of TCP connections, one Connection corresponds to one connection
// (用于处理Tcp连接的读写业务 一个连接对应一个Connection)
type Connection struct {
	// // The socket TCP socket of the current connection(当前连接的socket TCP套接字)
	conn net.Conn

	// The ID of the current connection, also known as SessionID, globally unique, used by server Connection
	// uint64 range: 0~18,446,744,073,709,551,615
	// This is the maximum number of connID theoretically supported by the process
	// (当前连接的ID 也可以称作为SessionID，ID全局唯一 ，服务端Connection使用
	// uint64 取值范围：0 ~ 18,446,744,073,709,551,615
	// 这个是理论支持的进程connID的最大数量)
	connID uint64

	// connection id for string
	// (字符串的连接id)
	connIdStr string

	// The workerID responsible for handling the link
	// 负责处理该连接的workerID
	workerID uint32

	// The message management module that manages MsgID and the corresponding processing method
	// (消息管理MsgID和对应处理方法的消息管理模块)
	msgHandler hinterface.IMsgHandle

	// Channel to notify that the connection has exited/stopped
	// (告知该连接已经退出/停止的channel)
	ctx    context.Context
	cancel context.CancelFunc

	// Go StartWriter Flag
	// (开始初始化写协程标志)
	startWriterFlag int32

	// Connection properties
	// (连接属性)
	property map[string]interface{}

	// Lock to protect the current property
	// (保护当前property的锁)
	propertyLock sync.Mutex

	// Which Connection Manager the current connection belongs to
	// (当前连接是属于哪个Connection Manager的)
	connManager hinterface.IConnManager

	// Connection name, default to be the same as the name of the Server/Client that created the connection
	// (连接名称，默认与创建连接的Server/Client的Name一致)
	name string
}

// newServerConn :for Server, method to create a Server-side connection with Server-specific properties
// (创建一个Server服务端特性的连接的方法)
func newServerConn(server hinterface.IServer, conn net.Conn, connID uint64) hinterface.IConnection {

	// Initialize Conn properties
	c := &Connection{
		conn:            conn,
		connID:          connID,
		connIdStr:       strconv.FormatUint(connID, 10),
		startWriterFlag: 0,
		property:        nil,
		name:            server.ServerName(),
	}
	// Inherited properties from server (从server继承过来的属性)
	c.msgHandler = server.GetMsgHandler()

	// Bind the current Connection with the Server's ConnManager
	// (将当前的Connection与Server的ConnManager绑定)
	c.connManager = server.GetConnMgr()

	// Add the newly created Conn to the connection manager
	// (将新创建的Conn添加到连接管理中)
	server.GetConnMgr().Add(c)

	return c
}

// newServerConn :for Server, method to create a Client-side connection with Client-specific properties
// (创建一个Client服务端特性的连接的方法)
func newClientConn(client hinterface.IClient, conn net.Conn) hinterface.IConnection {
	c := &Connection{
		conn:            conn,
		connID:          0,  // client ignore
		connIdStr:       "", // client ignore
		startWriterFlag: 0,
		property:        nil,
		name:            client.GetName(),
	}

	// Inherited properties from server (从client继承过来的属性)
	c.msgHandler = client.GetMsgHandler()

	return c
}

// StartWriter is the goroutine that writes messages to the client
// (写消息Goroutine， 用户将数据发送给客户端)
func (c *Connection) StartWriter() {
	zlog.Infof("Writer Goroutine is running")
	ticker := time.NewTicker(10 * time.Millisecond)
	defer func() {
		zlog.Infof("%s [conn Writer exit!]", c.RemoteAddr().String())
		ticker.Stop()
		c.Flush()
	}()
	for {
		select {
		case <-ticker.C:
			err := c.Flush()
			if err != nil {
				zlog.Errorf("Flush Buff Data error: %v Conn Writer exit", err)
				return
			}
		case <-c.ctx.Done():
			return
		}
	}
}

// StartReader is a goroutine that reads data from the client
// (读消息Goroutine，用于从客户端中读取数据)
func (c *Connection) StartReader() {
	zlog.Infof("[Reader Goroutine is running]")
	defer zlog.Infof("%s [conn Reader exit!]", c.RemoteAddr().String())
	defer c.Stop()
	defer func() {
		if err := recover(); err != nil {
			zlog.Errorf("connID=%d, panic err=%v", c.GetConnID(), err)
		}
	}()

	//Reduce buffer allocation times to improve efficiency
	// add by ray 2023-02-03
	buffer := make([]byte, hglobal.Config.Hin.IOReadBuffSize)

	for {
		select {
		case <-c.ctx.Done():
			return
		default:

			// read data from the connection's IO into the memory buffer
			// (从conn的IO中读取数据到内存缓冲buffer中)
			n, err := c.conn.Read(buffer)
			if err != nil {
				zlog.Errorf("read msg head [read datalen=%d], error = %s", n, err)
				return
			}
			zlog.Debugf("read buffer %s \n", hex.EncodeToString(buffer[0:n]))

		}
	}
}

// Start starts the connection and makes the current connection work.
// (启动连接，让当前连接开始工作)
func (c *Connection) Start() {
	defer func() {
		if err := recover(); err != nil {
			zlog.Errorf("Connection Start() error: %v", err)
		}
	}()
	c.ctx, c.cancel = context.WithCancel(context.Background())

	// 占用workerid
	c.workerID = useWorker(c)

	// Start the Goroutine for reading data from the client
	// (开启用户从客户端读取数据流程的Goroutine)
	go c.StartReader()

	select {
	case <-c.ctx.Done():
		c.finalizer()

		// 归还workerid
		freeWorker(c)
		return
	}
}

// Stop stops the connection and ends the current connection state.
// (停止连接，结束当前连接状态)
func (c *Connection) Stop() {
	c.cancel()
}

func (c *Connection) GetConnection() net.Conn {
	return c.conn
}

func (c *Connection) GetWsConn() *websocket.Conn {
	return nil
}

// Deprecated: use GetConnection instead
func (c *Connection) GetTCPConnection() net.Conn {
	return c.conn
}

func (c *Connection) GetConnID() uint64 {
	return c.connID
}

func (c *Connection) GetConnIdStr() string {
	return c.connIdStr
}

func (c *Connection) GetWorkerID() uint32 {
	return c.workerID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Connection) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *Connection) Send(data []byte) error {
	if c.isClosed() == true {
		return errors.New("connection closed when send msg")
	}
	_, err := c.conn.Write(data)
	if err != nil {
		zlog.Errorf("SendMsg err data = %+v, err = %+v", data, err)
		return err
	}
	return nil
}

// SendMsg directly sends Message data to the remote TCP client.
// (直接将Message数据发送数据给远程的TCP客户端)
func (c *Connection) SendMsg(msgID uint32, data []byte) error {

	if c.isClosed() == true {
		return errors.New("connection closed when send msg")
	}
	// Pack data and send it
	msg, err := c.packet.Pack(zpack.NewMsgPackage(msgID, data))
	if err != nil {
		zlog.Errorf("Pack error msg ID = %d", msgID)
		return errors.New("Pack error msg ")
	}

	err = c.Send(msg)
	if err != nil {
		zlog.Errorf("SendMsg err msg ID = %d, data = %+v, err = %+v", msgID, string(msg), err)
		return err
	}

	return nil
}

func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	if c.property == nil {
		c.property = make(map[string]interface{})
	}

	c.property[key] = value
}

func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	if value, ok := c.property[key]; ok {
		return value, nil
	}

	return nil, errors.New("no property found")
}

func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}

func (c *Connection) Context() context.Context {
	return c.ctx
}

func (c *Connection) finalizer() {
	// Close the socket connection
	_ = c.conn.Close()

	// Remove the connection from the connection manager
	if c.connManager != nil {
		c.connManager.Remove(c)
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				zlog.Errorf("Conn finalizer panic: %v", err)
			}
		}()

	}()

	zlog.Infof("Conn Stop()...ConnID = %d", c.connID)
}

func (c *Connection) GetName() string {
	return c.name
}

func (c *Connection) GetMsgHandler() hinterface.IMsgHandle {
	return c.msgHandler
}

func (c *Connection) isClosed() bool {
	return c.ctx == nil || c.ctx.Err() != nil
}

func (c *Connection) setStartWriterFlag() bool {
	return atomic.CompareAndSwapInt32(&c.startWriterFlag, 0, 1)
}
