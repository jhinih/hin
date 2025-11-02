package main

import (
	"fmt"
	"github.com/jhinih/hin/hinterface"
	"github.com/jhinih/hin/hnet"
	"strconv"
)

type Router1 struct {
	hnet.BaseRouter
}
type Router struct {
	hnet.BaseRouter
}

func (r *Router1) Handle(request hinterface.IRequest) {
	fmt.Println("handler1 run")
	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
	err := request.GetConnection().Send(201, []byte("Handle%%%%%%%%"))
	if err != nil {
		fmt.Println(err)
	}
}
func (r *Router) Handle(request hinterface.IRequest) {
	fmt.Println("handler0 run")
	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
	err := request.GetConnection().Send(200, []byte("Handle@@@@@@@@@"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := hnet.NewServer()
	s.AddRouter(1, &Router1{})
	s.AddRouter(0, &Router{})
	s.Server()
}
