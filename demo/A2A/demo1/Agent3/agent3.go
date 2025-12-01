package main

import (
	"fmt"
	"github.com/jhinih/hin/hinterface"
	"github.com/jhinih/hin/hnet"
	"strconv"
)

type Agent3Router1 struct {
	hnet.BaseRouter
}
type Agent3Router2 struct {
	hnet.BaseRouter
}

type Agent3Client1Router1 struct {
	hnet.BaseRouter
}

type Agent3Client1Router2 struct {
	hnet.BaseRouter
}

//	func (r *Router1) PreHandle(request hinterface.IRequest) {
//		fmt.Println("[PreHandler 1 run]")
//		fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
//		err := request.GetConnection().Send(201, []byte("PreHandle%%%%%%%%"))
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
func (r *Agent3Router1) Handle(request hinterface.IRequest) {
	fmt.Println("[Handler 1 run]")
	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
	err := request.GetConnection().Send(201, []byte("Handle%%%%%%%%"))
	if err != nil {
		fmt.Println(err)
	}
}

//func (r *Router1) PostHandle(request hinterface.IRequest) {
//	fmt.Println("[PostHandler 1 run]")
//	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
//	err := request.GetConnection().Send(201, []byte("PostHandle%%%%%%%%"))
//	if err != nil {
//		fmt.Println(err)
//	}
//}

//	func (r *Router) PreHandle(request hinterface.IRequest) {
//		fmt.Println("[PreHandler 0 run]")
//		fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
//		err := request.GetConnection().Send(200, []byte("PreHandle@@@@@@@@@"))
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
func (r *Agent3Router2) Handle(request hinterface.IRequest) {
	fmt.Println("[Handler 0 run]")
	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
	err := request.GetConnection().Send(200, []byte("Handle@@@@@@@@@"))
	if err != nil {
		fmt.Println(err)
	}
}

//func (r *Router) PostHandle(request hinterface.IRequest) {
//	fmt.Println("[PostHandler 0 run]")
//	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
//	err := request.GetConnection().Send(200, []byte("PostHandle@@@@@@@@@"))
//	if err != nil {
//		fmt.Println(err)
//	}
//}

func (r *Agent3Client1Router1) Handle(request hinterface.IRequest) {
	fmt.Println("[Handler 0 run]")
	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
	err := request.GetConnection().Send(200, []byte("Handle@@@@@@@@@"))
	if err != nil {
		fmt.Println(err)
	}
}

//func (r *Router) PostHandle(request hinterface.IRequest) {
//	fmt.Println("[PostHandler 0 run]")
//	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
//	err := request.GetConnection().Send(200, []byte("PostHandle@@@@@@@@@"))
//	if err != nil {
//		fmt.Println(err)
//	}
//}

//	func (r *Router) PreHandle(request hinterface.IRequest) {
//		fmt.Println("[PreHandler 0 run]")
//		fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
//		err := request.GetConnection().Send(200, []byte("PreHandle@@@@@@@@@"))
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
func (r *Agent3Client1Router2) Handle(request hinterface.IRequest) {
	fmt.Println("[Handler 0 run]")
	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
	err := request.GetConnection().Send(200, []byte("Handle@@@@@@@@@"))
	if err != nil {
		fmt.Println(err)
	}
}

func Agent3Server1ConnectionStartHook(connection hinterface.IConnection) {
	fmt.Println("=====》[connection", strconv.Itoa(int(connection.GetConnectionID())), "start]")
	connection.SetProperty("name", "jhinih")
	if err := connection.Send(201, []byte("huh huh")); err != nil {
		fmt.Println(err)
	}
}

func Agent3Server1ConnectionStopHook(connection hinterface.IConnection) {
	fmt.Println("=====》[connection", strconv.Itoa(int(connection.GetConnectionID())), "stop]")
	fmt.Println(connection.GetProperty("name"))
	if err := connection.Send(201, []byte("bye bye")); err != nil {
		fmt.Println(err)
	}
}

func Agent3Server2ConnectionStartHook(connection hinterface.IConnection) {
	fmt.Println("=====》[connection", strconv.Itoa(int(connection.GetConnectionID())), "start]")
	connection.SetProperty("name", "jhinih")
	if err := connection.Send(201, []byte("huh huh")); err != nil {
		fmt.Println(err)
	}
}

func Agent3Server2ConnectionStopHook(connection hinterface.IConnection) {
	fmt.Println("=====》[connection", strconv.Itoa(int(connection.GetConnectionID())), "stop]")
	fmt.Println(connection.GetProperty("name"))
	if err := connection.Send(201, []byte("bye bye")); err != nil {
		fmt.Println(err)
	}
}

func Agent3Client1ConnectionStartHook(connection hinterface.IConnection) {
	fmt.Println("=====》[connection", strconv.Itoa(int(connection.GetConnectionID())), "start]")
	connection.SetProperty("name", "jhinih")
	if err := connection.Send(201, []byte("huh huh")); err != nil {
		fmt.Println(err)
	}
}

func Agent3Client1ConnectionStopHook(connection hinterface.IConnection) {
	fmt.Println("=====》[connection", strconv.Itoa(int(connection.GetConnectionID())), "stop]")
	fmt.Println(connection.GetProperty("name"))
	if err := connection.Send(201, []byte("bye bye")); err != nil {
		fmt.Println(err)
	}
}

func main() {
	a := hnet.NewAgent()
	a.SetIdentity("ID", 1)
	a.SetIdentity("Name", "王鑫宇")
	a.SetIdentity("Position", "预言家")

	s1 := hnet.NewServer()
	s1.SetConnectionStartHook(Agent3Server1ConnectionStartHook)
	s1.SetConnectionStopHook(Agent3Server1ConnectionStopHook)
	s1.AddRouter(1, &Agent3Router1{})

	s2 := hnet.NewServer()
	s2.SetConnectionStartHook(Agent3Server2ConnectionStartHook)
	s2.SetConnectionStopHook(Agent3Server2ConnectionStopHook)
	s2.AddRouter(1, &Agent3Router2{})

	a.AddServer(0, s1)
	a.AddServer(1, s2)

	c1 := hnet.NewClient("192.168.1.90", 8080)
	c1.SetConnectionStartHook(Agent3Client1ConnectionStartHook)
	c1.SetConnectionStopHook(Agent3Client1ConnectionStopHook)
	c1.AddRouter(1, &Agent3Client1Router1{})
	//c1.AddRouter(1, &Agent1Client1Router2{})
	a.Start()
}
