package hnet

import (
	"github.com/jhinih/hin/hconfig"
	"github.com/jhinih/hin/hinterface"
)

type Server struct {
	Name string

	IPVersion string
	IP        string
	Port      int
}

func newServerWithConfig(config *hconfig.Config, ipVersion string, opts ...Option) hinterface.IServer {
	s := &Server{
		Name:      config.Server.Name,
		IPVersion: ipVersion,
		IP:        config.Server.Host,
		Port:      config.Server.TCPPort,
	}

	return s
}

// NewServer creates a server handle
// (创建一个服务器句柄)
func NewServer(opts ...Option) hinterface.IServer {
	return newServerWithConfig(hconfig.Conf, "tcp", opts...)
}
func (s *Server) Start()  {}
func (s *Server) Stop()   {}
func (s *Server) Server() {}
