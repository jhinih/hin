package hnet

import (
	"fmt"
	"github.com/jhinih/hin/hinterface"
	"github.com/jhinih/hin/hpack"
	"net"
)

type Server struct {
	Name string
	IP   string
	Port int

	IPVersion         string
	MsgHandler        hinterface.IMessageHandler
	ConnectionManager hinterface.IConnectionManager

	exitChan            chan any
	ConnectionStartHook func(hinterface.IConnection)
	ConnectionStopHook  func(hinterface.IConnection)

	Pack hinterface.IPack
}

func NewServer() hinterface.IServer {
	return &Server{
		Name:      "HinServer",
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,

		MsgHandler:        NewServerMessageHandler(),
		ConnectionManager: NewServerConnectionManager(),

		Pack: hpack.NewLTVPack(),
	}
}

func (s *Server) Start() {
	fmt.Println("[server start]")

	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println(err)
	}
	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println(err)
	}
	cid := uint32(0)
	go func() {
		s.MsgHandler.StartWorkPoll()
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println(err)
			}

			if s.ConnectionManager.Len() >= 100 /*最大连接数*/ {
				conn.Close()
				continue
			}
			dealConn := NewServerConnection(s, conn, cid, s.MsgHandler)
			cid++

			go dealConn.Start()
		}
	}()

	select {}
}
func (s *Server) Stop() {
	s.ConnectionManager.ClearConnection()
	s.exitChan <- struct{}{}
	close(s.exitChan)
}
func (s *Server) Serve() {
	s.Start()

	select {}
}

func (s *Server) AddRouter(msgID uint32, router hinterface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
}

func (s *Server) SetConnectionStartHook(fn func(hinterface.IConnection)) {
	s.ConnectionStartHook = fn
}
func (s *Server) SetConnectionStopHook(fn func(hinterface.IConnection)) {
	s.ConnectionStopHook = fn
}
func (s *Server) GetPack() hinterface.IPack {
	return s.Pack
}
func (s *Server) SetPack(pack hinterface.IPack) {
	s.Pack = pack
}
func (s *Server) GetConnectionStartHook() func(hinterface.IConnection) {
	return s.ConnectionStartHook
}
func (s *Server) GetConnectionStopHook() func(hinterface.IConnection) {
	return s.ConnectionStopHook
}
func (s *Server) GetMsgHandler() hinterface.IMessageHandler {
	return s.MsgHandler
}
func (s *Server) GetConnectionManager() hinterface.IConnectionManager {
	return s.ConnectionManager
}
func (s *Server) GetName() string {
	return s.Name
}

func init() {}
