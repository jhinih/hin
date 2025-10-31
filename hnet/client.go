package hnet

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/jhinih/hin/hinterface"
	"github.com/jhinih/hin/hlog/zlog"
	"net"
	"net/http"
	"net/url"
	"sync"
)

type Client struct {
	sync.WaitGroup
	sync.Mutex
	started bool
	ctx     context.Context
	cancel  context.CancelFunc
	// Client Name 客户端的名称
	Name string
	// IP of the target server to connect 目标连接服务器的IP
	Ip string
	// Port of the target server to connect 目标连接服务器的端口
	Port int
	Url  *url.URL // 扩展，连接时带上其他参数
	// Custom headers for WebSocket connection WebSocket连接的自定义头信息
	WsHeader http.Header
	// Client version tcp,websocket,客户端版本 tcp,websocket
	version string
	// Connection instance 连接实例
	conn hinterface.IConnection
	// Connection instance 连接实例的锁，保证可见性
	connMux sync.Mutex
	// Hook function called on connection start 该client的连接创建时Hook函数
	onConnStart func(conn hinterface.IConnection)
	// Hook function called on connection stop 该client的连接断开时的Hook函数
	onConnStop func(conn hinterface.IConnection)
	// Asynchronous channel for capturing connection close status 异步捕获连接关闭状态
	// exitChan chan struct{}
	// Message management module 消息管理模块
	msgHandler hinterface.IMsgHandle
	// Use TLS 使用TLS
	useTLS bool
	// Error channel
	errChan chan error
}

func NewClient(ip string, port int, opts ...ClientOption) hinterface.IClient {

	c := &Client{
		// Default name, can be modified using the WithNameClient Option
		// (默认名称，可以使用WithNameClient的Option修改)
		Name:    "ZinxClientTcp",
		Ip:      ip,
		Port:    port,
		version: "tcp",
		errChan: make(chan error, 1),
	}
	return c
}

// notify error unblock
func (c *Client) notifyErr(err error) {
	select {
	case c.errChan <- err:
	default:
	}
}

// Start starts the client, sends requests and establishes a connection.
// (重新启动客户端，发送请求且建立连接)
func (c *Client) Restart() {
	//try to stop and wait until client stoped
	c.Stop()

	//set started flag
	c.Lock()
	if c.started {
		// already started, just return
		c.Unlock()
		return
	}
	c.started = true
	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.Add(1)
	c.Unlock()

	zlog.Infof("[START] Zinx Client dial RemoteAddr: %s:%d\n", c.Ip, c.Port)
	go func() {
		defer c.Done()

		// Create a raw socket and get net.Conn (创建原始Socket，得到net.Conn)
		var connect hinterface.IConnection

		var conn net.Conn
		var err error
		if c.useTLS {
			// TLS encryption
			config := &tls.Config{
				// Skip certificate verification here because the CA certificate of the certificate issuer is not authenticated
				// (这里是跳过证书验证，因为证书签发机构的CA证书是不被认证的)
				InsecureSkipVerify: true,
			}
			d := &tls.Dialer{
				Config: config,
			}
			//conn, err = tls.Dial("tcp", fmt.Sprintf("%v:%v", net.ParseIP(c.Ip), c.Port), config)
			conn, err = d.DialContext(c.ctx, "tcp", fmt.Sprintf("%v:%v", net.ParseIP(c.Ip), c.Port))
			if err != nil {
				zlog.Errorf("tls client connect to server failed, err:%v", err)
				c.notifyErr(err)
				return
			}
		} else {
			//conn, err = net.DialTCP("tcp", nil, addr)
			d := &net.Dialer{}
			conn, err = d.DialContext(c.ctx, "tcp", fmt.Sprintf("%v:%v", net.ParseIP(c.Ip), c.Port))
			if err != nil {
				// connection failed
				zlog.Errorf("client connect to server failed, err:%v", err)
				c.notifyErr(err)
				return
			}
		}
		// Create Connection object
		connect = newClientConn(c, conn)

		// Set connection to the client
		c.setConn(connect)
		// Start connection
		go connect.Start()

		<-c.ctx.Done()
		zlog.Infof("client exit.")
	}()
}

// Start starts the client, sends requests and establishes a connection.
// (启动客户端，发送请求且建立连接)
func (c *Client) Start() {
	c.Restart()
}

// 保证重复调用Stop不会导致panic
func (c *Client) Stop() {
	c.Lock()
	defer c.Unlock()
	if !c.started {
		return
	}
	c.started = false

	con := c.Conn()
	if con != nil {
		con.Stop()
	}

	// c.exitChan <- struct{}{}
	// close(c.exitChan)
	// close(c.ErrChan)
	if c.cancel != nil {
		c.cancel()
	}
	c.Wait()
}

func (c *Client) AddRouter(msgID uint32, router hinterface.IRouter) {
	c.msgHandler.AddRouter(msgID, router)
}

func (c *Client) Conn() hinterface.IConnection {
	c.connMux.Lock()
	defer c.connMux.Unlock()
	return c.conn
}

func (c *Client) setConn(con hinterface.IConnection) {
	c.connMux.Lock()
	defer c.connMux.Unlock()
	c.conn = con
}

func (c *Client) SetOnConnStart(hookFunc func(hinterface.IConnection)) {
	c.onConnStart = hookFunc
}

func (c *Client) SetOnConnStop(hookFunc func(hinterface.IConnection)) {
	c.onConnStop = hookFunc
}

func (c *Client) GetOnConnStart() func(hinterface.IConnection) {
	return c.onConnStart
}

func (c *Client) GetOnConnStop() func(hinterface.IConnection) {
	return c.onConnStop
}

func (c *Client) GetMsgHandler() hinterface.IMsgHandle {
	return c.msgHandler
}

func (c *Client) GetErrChan() <-chan error {
	return c.errChan
}

func (c *Client) SetName(name string) {
	c.Name = name
}

func (c *Client) GetName() string {
	return c.Name
}

func (c *Client) SetUrl(url *url.URL) {
	c.Url = url
}

func (c *Client) GetUrl() *url.URL {
	return c.Url
}

func (c *Client) SetWsHeader(header http.Header) {
	c.WsHeader = header
}

func (c *Client) GetWsHeader() http.Header {
	return c.WsHeader
}
