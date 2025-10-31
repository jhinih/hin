package hnet

import (
	"errors"
	"fmt"
	"github.com/jhinih/hin/hconfig"
	"github.com/jhinih/hin/hinterface"
	"github.com/jhinih/hin/hlog/zlog"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int

	// Current server's connection manager (当前Server的连接管理器)
	ConnMgr hinterface.IConnManager
	// connection id
	cID uint64

	// Current server's message handler module, used to bind MsgID to corresponding processing methods
	// (当前Server的消息管理模块，用来绑定MsgID和对应的处理方法)
	msgHandler hinterface.IMsgHandle
}

func newServerWithConfig(config *hconfig.Config, ipVersion string, opts ...Option) hinterface.IServer {
	s := &Server{
		Name:      "config.Server.Name,",
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}

	return s
}

// NewServer creates a server handle
// (创建一个服务器句柄)
func NewServer(opts ...Option) hinterface.IServer {
	return newServerWithConfig(hconfig.Conf, "tcp", opts...)
}
func (s *Server) Start() {
	zlog.Infof("[START] TCP Server name: %s,listener at IP: %s, Port %d is starting", s.Name, s.IP, s.Port)
	// 1. Get a TCP address
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		zlog.Errorf("[START] resolve tcp addr err: %v\n", err)
		return
	}

	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		zlog.Errorf("[START] listen tcp err: %v\n", err)
	}

	// 3. Start server network connection business
	go func() {
		for {
			// 3.2 Block and wait for a client to establish a connection request.
			// (阻塞等待客户端建立连接请求)
			conn, err := listener.Accept()
			if err != nil {
				if errors.Is(err, net.ErrClosed) {
					zlog.Errorf("Listener closed")
					return
				}
				zlog.Errorf("Accept err: %v", err)
				continue
			}

			go func() {
				for {
					readbuf := make([]byte, 1024)
					n, err := conn.Read(readbuf)
					if err != nil {
						zlog.Errorf("Read err: %v", err)
						continue
					}
					if _, err = conn.Write(readbuf[:n]); err != nil {
						continue
					}
				}
			}()
		}
	}()
	//select {
	//case <-s.exitChan:
	//	err := listener.Close()
	//	if err != nil {
	//		zlog.Ins().ErrorF("listener close err: %v", err)
	//	}
	//}
}
func (s *Server) Stop() {
	//TODO
}
func (s *Server) Server() {
	s.Start()

	//TODO

	select {}
}

//func (s *Server) GetConnMgr() IConnManager {
//
//}
//
//func (s *Server) AddRouter(msgID uint32, router hinterface.IRouter) {
//
//}
//
//func (s *Server) GetMsgHandler() IMsgHandle {
//
//}
