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
	fmt.Println(string(request.GetMsgData()))
	err := request.GetConnection().Send(1, []byte("PreHandle&&"))
	if err != nil {
		fmt.Println(err)
	}
}
func (r *Router) Handle(request hinterface.IRequest) {
	fmt.Println("[Handle]")
	err := request.GetConnection().Send(1, []byte("Handle%%"))
	if err != nil {
		fmt.Println(err)
	}
}
func (r *Router) PostHandle(request hinterface.IRequest) {
	fmt.Println("[PostHandle]")
	err := request.GetConnection().Send(1, []byte("PostHandle^^"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := hnet.NewServer()
	s.AddRouter(&Router{})
	s.Server()
}
