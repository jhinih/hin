package hnet

import (
	"context"
	"fmt"
	"github.com/jhinih/hin/hinitialize"
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

	IPVersion  string
	MsgHandler hinterface.IMessageHandler

	ConnectionStartHook func(hinterface.IConnection)
	ConnectionStopHook  func(hinterface.IConnection)

	Pack          hinterface.IPack
	connection    hinterface.IConnection
	connectionMux sync.Mutex

	errChan chan error
}

func NewClient(ip string, port int, opts ...ClientOption) hinterface.IClient {
	hinitialize.Init()
	c := &Client{
		Name:       "HinClient",
		IP:         ip,
		Port:       port,
		IPVersion:  "tcp4",
		MsgHandler: NewClientMessageHandler(),
		Pack:       hpack.NewTLVPack(),
		errChan:    make(chan error, 1),
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
		c.MsgHandler.StartWorkPoll()
		defer c.Done()
		d := &net.Dialer{}
		conn, err := d.DialContext(c.ctx, "tcp", fmt.Sprintf("%v:%v", net.ParseIP(c.IP), c.Port))
		if err != nil {
			fmt.Println("connection error", err)
			c.notifyErr(err)
			return
		}
		connection := NewClientConnection(c, conn)
		c.SetConnection(connection)

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

	connection := c.GetConnection()
	if connection != nil {
		fmt.Println("[Client stop]")
		connection.Stop()
		hinitialize.Eve()
	}

	if c.cancel != nil {
		c.cancel()
	}
	c.Wait()
}
func (c *Client) AddRouter(msgID uint32, router hinterface.IRouter) {
	c.MsgHandler.AddRouter(msgID, router)
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
	return c.MsgHandler
}
func (c *Client) GetConnection() hinterface.IConnection {
	c.connectionMux.Lock()
	defer c.connectionMux.Unlock()
	return c.connection
}
func (c *Client) SetConnection(connection hinterface.IConnection) {
	c.connectionMux.Lock()
	defer c.connectionMux.Unlock()
	c.connection = connection
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
