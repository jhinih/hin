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

func (r *Router1) PreHandle(request hinterface.IRequest) {
	fmt.Println("[PreHandler 1 run]")
	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
	err := request.GetConnection().Send(201, []byte("PreHandle%%%%%%%%"))
	if err != nil {
		fmt.Println(err)
	}
}
func (r *Router1) Handle(request hinterface.IRequest) {
	fmt.Println("[Handler 1 run]")
	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
	err := request.GetConnection().Send(201, []byte("Handle%%%%%%%%"))
	if err != nil {
		fmt.Println(err)
	}
}
func (r *Router1) PostHandle(request hinterface.IRequest) {
	fmt.Println("[PostHandler 1 run]")
	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
	err := request.GetConnection().Send(201, []byte("PostHandle%%%%%%%%"))
	if err != nil {
		fmt.Println(err)
	}
}
func (r *Router) PreHandle(request hinterface.IRequest) {
	fmt.Println("[PreHandler 0 run]")
	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
	err := request.GetConnection().Send(200, []byte("PreHandle@@@@@@@@@"))
	if err != nil {
		fmt.Println(err)
	}
}
func (r *Router) Handle(request hinterface.IRequest) {
	fmt.Println("[Handler 0 run]")
	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
	err := request.GetConnection().Send(200, []byte("Handle@@@@@@@@@"))
	if err != nil {
		fmt.Println(err)
	}
}
func (r *Router) PostHandle(request hinterface.IRequest) {
	fmt.Println("[PostHandler 0 run]")
	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
	err := request.GetConnection().Send(200, []byte("PostHandle@@@@@@@@@"))
	if err != nil {
		fmt.Println(err)
	}
}

func ConnectionStartHook(connection hinterface.IConnection) {
	fmt.Println("=====》[connection", strconv.Itoa(int(connection.GetConnectionID())), "start]")
	connection.SetProperty("name", "jhinih")
	if err := connection.Send(114514, []byte("huh huh")); err != nil {
		fmt.Println(err)
	}
}

func ConnectionStopHook(connection hinterface.IConnection) {
	fmt.Println("=====》[connection", strconv.Itoa(int(connection.GetConnectionID())), "stop]")
	fmt.Println(connection.GetProperty("name"))
	if err := connection.Send(404, []byte("bye bye")); err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := hnet.NewServer()
	s.SetConnectionStartHook(ConnectionStartHook)
	s.SetConnectionStopHook(ConnectionStopHook)
	s.AddRouter(1, &Router1{})
	s.AddRouter(0, &Router{})
	s.Serve()
}
