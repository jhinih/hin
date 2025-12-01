package main

import (
	"fmt"
	"github.com/jhinih/hin/hinterface"
	"github.com/jhinih/hin/hnet"
	"github.com/jhinih/hin/hpack"
	"strconv"
)

type Agent1Server1Router1 struct {
	hnet.BaseRouter
}
type Agent1Server2Router1 struct {
	hnet.BaseRouter
}
type Agent1Server2Router2 struct {
	hnet.BaseRouter
}

type Agent1Client1Router1 struct {
	hnet.BaseRouter
}

type Agent1Client1Router2 struct {
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
func (r *Agent1Server1Router1) Handle(request hinterface.IRequest) {
	fmt.Println("[Handler 1 run]")
	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
	err := request.GetConnection().Send(201, []byte("Handle%%%%%%%%"))
	if err != nil {
		fmt.Println(err)
	}

	c1 := hnet.NewClient("127.168.1.90", 8080)
	defer c1.Stop()
	c1.Start()
	msg := &hpack.Message{
		ID:   1,
		Data: []byte("Hello World"),
	}
	request1 := hnet.NewRequest(c1.GetConnection(), msg)
	c1.GetMsgHandler().SendMsg2TaskQueue(request1)

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
func (r *Agent1Server2Router1) Handle(request hinterface.IRequest) {
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
func (r *Agent1Server2Router2) Handle(request hinterface.IRequest) {
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
func (r *Agent1Client1Router1) Handle(request hinterface.IRequest) {
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
func (r *Agent1Client1Router2) Handle(request hinterface.IRequest) {
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

func Agent1Server1ConnectionStartHook(connection hinterface.IConnection) {
	fmt.Println("=====》[connection", strconv.Itoa(int(connection.GetConnectionID())), "start]")
	connection.SetProperty("name", "jhinih")
	if err := connection.Send(201, []byte("huh huh")); err != nil {
		fmt.Println(err)
	}
}

func Agent1Server1ConnectionStopHook(connection hinterface.IConnection) {
	fmt.Println("=====》[connection", strconv.Itoa(int(connection.GetConnectionID())), "stop]")
	fmt.Println(connection.GetProperty("name"))
	if err := connection.Send(201, []byte("bye bye")); err != nil {
		fmt.Println(err)
	}
}

func Agent1Server2ConnectionStartHook(connection hinterface.IConnection) {
	fmt.Println("=====》[connection", strconv.Itoa(int(connection.GetConnectionID())), "start]")
	connection.SetProperty("name", "jhinih")
	if err := connection.Send(201, []byte("huh huh")); err != nil {
		fmt.Println(err)
	}
}

func Agent1Server2ConnectionStopHook(connection hinterface.IConnection) {
	fmt.Println("=====》[connection", strconv.Itoa(int(connection.GetConnectionID())), "stop]")
	fmt.Println(connection.GetProperty("name"))
	if err := connection.Send(201, []byte("bye bye")); err != nil {
		fmt.Println(err)
	}
}

func Agent1Client1ConnectionStartHook(connection hinterface.IConnection) {
	fmt.Println("=====》[connection", strconv.Itoa(int(connection.GetConnectionID())), "start]")
	connection.SetProperty("name", "jhinih")
	if err := connection.Send(201, []byte("huh huh")); err != nil {
		fmt.Println(err)
	}
}

func Agent1Client1ConnectionStopHook(connection hinterface.IConnection) {
	fmt.Println("=====》[connection", strconv.Itoa(int(connection.GetConnectionID())), "stop]")
	fmt.Println(connection.GetProperty("name"))
	if err := connection.Send(201, []byte("bye bye")); err != nil {
		fmt.Println(err)
	}
}

func main() {
	a := hnet.NewAgent()
	a.SetIdentity("ID", 2)
	a.SetIdentity("Name", "wxy")
	a.SetIdentity("Position", "狼人")

	s1 := hnet.NewServer()
	s1.SetConnectionStartHook(Agent1Server1ConnectionStartHook)
	s1.SetConnectionStopHook(Agent1Server1ConnectionStopHook)
	s1.AddRouter(1, &Agent1Server1Router1{})

	s2 := hnet.NewServer()
	s2.SetConnectionStartHook(Agent1Server2ConnectionStartHook)
	s2.SetConnectionStopHook(Agent1Server2ConnectionStopHook)
	s2.AddRouter(1, &Agent1Server2Router1{})
	s2.AddRouter(2, &Agent1Server2Router2{})

	a.AddServer(0, s1)
	a.AddServer(1, s2)

	c1 := hnet.NewClient("192.168.1.90", 8080)
	c1.SetConnectionStartHook(Agent1Client1ConnectionStartHook)
	c1.SetConnectionStopHook(Agent1Client1ConnectionStopHook)
	c1.AddRouter(1, &Agent1Client1Router1{})
	//c1.AddRouter(1, &Agent1Client1Router2{})
	a.Start()
}
