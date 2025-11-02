package hnet

import (
	"fmt"
	"github.com/jhinih/hin/hinterface"
	"net"
)

type Server struct {
	Name string
	IP   string
	Port int

	IPVersion         string
	MsgHandler        hinterface.IMessageHandler
	ConnectionManager hinterface.IConnectionManager

	exitChan        chan any
	ConnectionStart func(hinterface.IConnection)
	ConnectionStop  func(hinterface.IConnection)
}

func NewServer() hinterface.IServer {
	return &Server{
		IPVersion:         "tcp4",
		IP:                "127.0.0.1",
		Port:              8999,
		Name:              "hnet",
		MsgHandler:        NewMessageHandler(),
		ConnectionManager: NewConnectionManager(),
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
			dealConn := NewConnection(s, conn, cid, s.MsgHandler)
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
func (s *Server) Server() {
	s.Start()

	select {}
}

func (s *Server) AddRouter(msgID uint32, router hinterface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
}

func (s *Server) GetConnectionManagerHandler() hinterface.IConnectionManager {
	return s.ConnectionManager
}

func (s *Server) SetConnectionStartHook(fn func(hinterface.IConnection)) {
	s.ConnectionStart = fn
}
func (s *Server) SetConnectionStopHook(fn func(hinterface.IConnection)) {
	s.ConnectionStop = fn
}
func (s *Server) CallConnectionStartHook(connection hinterface.IConnection) {
	if s.ConnectionStart != nil {
		fmt.Println("call——————>connection start hook")
		s.ConnectionStart(connection)
	}

}
func (s *Server) CallConnectionStopHook(connection hinterface.IConnection) {
	if s.ConnectionStop != nil {
		fmt.Println("call——————>connection stop hook")
		s.ConnectionStop(connection)
	}

}
