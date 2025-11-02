package main

import (
	"fmt"
	"github.com/jhinih/hin/hinterface"
	"github.com/jhinih/hin/hnet"
)

type Router struct {
	hnet.BaseRouter
}

func (r *Router) PreHandle(request hinterface.IRequest) {
	fmt.Println("[PreHandle]")
	_, err := request.GetConnection().GetTCpConnection().Write([]byte("PreHandle \n"))
	if err != nil {
		fmt.Println(err)
	}
}
func (r *Router) Handle(request hinterface.IRequest) {
	fmt.Println("[Handle]")
	_, err := request.GetConnection().GetTCpConnection().Write([]byte("Handle \n"))
	if err != nil {
		fmt.Println(err)
	}
}
func (r *Router) PostHandle(request hinterface.IRequest) {
	fmt.Println("[PostHandle]")
	_, err := request.GetConnection().GetTCpConnection().Write([]byte("PostHandle \n"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := hnet.NewServer()
	s.AddRouter(&Router{})
	s.Server()
}
