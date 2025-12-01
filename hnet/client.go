package hnet

import (
	"context"
	"fmt"
	"github.com/jhinih/hin/hinitialize"
	"strings"
	"sync"

	"github.com/jhinih/hin/hinterface"
	"github.com/jhinih/hin/hpack"
	"net"
)

type Client struct {
	sync.WaitGroup
	sync.Mutex
	started bool
	ctx     context.Context
	cancel  context.CancelFunc

	Name string
	IP   string
	Port int

	IPVersion         string
	MessageHandler    hinterface.IMessageHandler
	ConnectionManager hinterface.IConnectionManager

	exitChan            chan any
	ConnectionStartHook func(hinterface.IConnection)
	ConnectionStopHook  func(hinterface.IConnection)

	Pack          hinterface.IPack
	connectionMux sync.Mutex

	errChan chan error
}

func NewClient(ip string, port int, opts ...ClientOption) hinterface.IClient {
	hinitialize.Init()
	c := &Client{
		Name:      "HinClient",
		IPVersion: "tcp4",
		IP:        ip,
		Port:      port,

		MessageHandler:    NewClientMessageHandler(),
		ConnectionManager: NewClientConnectionManager(),

		exitChan: make(chan any),
		Pack:     hpack.NewTLVPack(),
		errChan:  make(chan error, 1),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
func (c *Client) notifyErr(err error) {
	select {
	case c.errChan <- err:
	default:
	}
}

func (c *Client) ReStart() {
	c.Stop()
	c.Lock()
	if c.started {
		c.Unlock()
		return
	}
	c.started = true
	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.Add(1)
	c.Unlock()
	go func() {
		c.MessageHandler.StartWorkPoll()
		defer c.Done()
		d := &net.Dialer{}
		conn, err := d.DialContext(c.ctx, "tcp", fmt.Sprintf("%v:%v", net.ParseIP(c.IP), c.Port))
		if err != nil {
			fmt.Println("connection error", err)
			c.notifyErr(err)
			return
		}
		connection := NewClientConnection(c, conn)
		c.AddConnection(connection)

		go connection.Start()
		<-c.ctx.Done()
	}()
	select {
	case _ = <-c.errChan:
		c.Lock()
		c.started = false
		c.Unlock()
	}

}
func (c *Client) Start() {
	fmt.Println("[Client start]")
	c.ReStart()
}
func (c *Client) Stop() {
	c.Lock()
	defer c.Unlock()
	if !c.started {
		return
	}
	c.started = false

	c.ConnectionManager.ClearConnection()
	hinitialize.Eve()

	if c.cancel != nil {
		c.cancel()
	}
	c.Wait()
}
func (c *Client) AddRouter(messageID uint32, router hinterface.IRouter) {
	c.MessageHandler.AddRouter(messageID, router)
}

func (c *Client) SetConnectionStartHook(fn func(hinterface.IConnection)) {
	c.ConnectionStartHook = fn
}
func (c *Client) SetConnectionStopHook(fn func(hinterface.IConnection)) {
	c.ConnectionStopHook = fn
}
func (c *Client) GetPack() hinterface.IPack {
	return c.Pack
}
func (c *Client) SetPack(pack hinterface.IPack) {
	c.Pack = pack
}
func (c *Client) GetConnectionStartHook() func(hinterface.IConnection) {
	return c.ConnectionStartHook
}
func (c *Client) GetConnectionStopHook() func(hinterface.IConnection) {
	return c.ConnectionStopHook
}
func (c *Client) GetMsgHandler() hinterface.IMessageHandler {
	return c.MessageHandler
}
func (c *Client) GetConnection(connectionID uint32) (hinterface.IConnection, error) {
	c.connectionMux.Lock()
	defer c.connectionMux.Unlock()
	connection, err := c.ConnectionManager.Get(connectionID)
	return connection, err

}
func (c *Client) AddConnection(connection hinterface.IConnection) {
	c.connectionMux.Lock()
	defer c.connectionMux.Unlock()
	c.ConnectionManager.Add(connection)
}

func (c *Client) SetName(name string) {
	c.Name = name
}
func (c *Client) GetName() string {
	return c.Name
}
func (c *Client) GetErrChan() <-chan error {
	return c.errChan
}

func (c *Client) GetClientID() uint64 {
	var b []byte
	for _, seg := range strings.Split(c.IP, ".") {
		b = append(b, fmt.Sprintf("%03s", seg)...)
	}
	b = append(b, fmt.Sprintf("%05d", c.Port)...)
	var ClientID uint64
	fmt.Sscanf(string(b), "%d", &ClientID)
	return ClientID
}

// // 将 12700000000108080 还原成 ip、port
//	func IDToIPPort(id uint64) (ip string, port int) {
//		// 取后 5 位 = 端口
//		port = int(id % 100000)
//		id /= 100000
//
//		// 从低位到高位依次取 3 位
//		segs := make([]string, 4)
//		for i := 3; i >= 0; i-- {
//			segs[i] = fmt.Sprintf("%03d", id%1000)
//			id /= 1000
//		}
//
//		ip = fmt.Sprintf("%s.%s.%s.%s", segs[0], segs[1], segs[2], segs[3])
//		return ip, port
//	}

//func (c *Client) SetUrl(url *url.URL) {
//	c.Url = url
//}
//
//func (c *Client) GetUrl() *url.URL {
//	return c.Url
//}
//
//func (c *Client) SetWsHeader(header http.Header) {
//	c.WsHeader = header
//}
//
//func (c *Client) GetWsHeader() http.Header {
//	return c.WsHeader
//}
