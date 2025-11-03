package main

import (
	"fmt"
	"github.com/jhinih/hin/hinterface"
	"github.com/jhinih/hin/hnet"
	"strconv"
	"time"
)

type ClientRouter1 struct {
	hnet.BaseRouter
}
type ClientRouter struct {
	hnet.BaseRouter
}

//	func (r *ClientRouter1) PreHandle(request hinterface.IRequest) {
//		fmt.Println("[Client PreHandler 1 run]")
//		fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
//		err := request.GetConnection().Send(1, []byte("PreHandle%%%%%%%%"))
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
func (r *ClientRouter1) Handle(request hinterface.IRequest) {
	fmt.Println("[Client Handler 201 run]")
	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
	err := request.GetConnection().Send(1, []byte("Handle%%%%%%%%"))
	if err != nil {
		fmt.Println(err)
	}
}

//	func (r *ClientRouter1) PostHandle(request hinterface.IRequest) {
//		fmt.Println("[Client PostHandler 1 run]")
//		fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
//		err := request.GetConnection().Send(1, []byte("PostHandle%%%%%%%%"))
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
//
//	func (r *ClientRouter) PreHandle(request hinterface.IRequest) {
//		fmt.Println("[Client PreHandler 0 run]")
//		fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
//		err := request.GetConnection().Send(0, []byte("PreHandle@@@@@@@@@"))
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
func (r *ClientRouter) Handle(request hinterface.IRequest) {
	fmt.Println("[Client Handler 200 run]")
	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
	err := request.GetConnection().Send(0, []byte("Handle@@@@@@@@@"))
	if err != nil {
		fmt.Println(err)
	}
}

//func (r *ClientRouter) PostHandle(request hinterface.IRequest) {
//	fmt.Println("[Client PostHandler 0 run]")
//	fmt.Println("Recv msg", strconv.Itoa(int(request.GetMsgID())), string(request.GetMsgData()))
//	err := request.GetConnection().Send(0, []byte("PostHandle@@@@@@@@@"))
//	if err != nil {
//		fmt.Println(err)
//	}
//}

func ClientConnectionStartHook(connection hinterface.IConnection) {
	fmt.Println("=====》[Client connection", strconv.Itoa(int(connection.GetConnectionID())), "start]")
	connection.SetProperty("web", "hin")
	for {
		if err := connection.Send(1, []byte("ClientConnectionStartHook send : {huh huh}")); err != nil {
			fmt.Println(err)
		}
		time.Sleep(5 * time.Second)
	}
}

func ClientConnectionStopHook(connection hinterface.IConnection) {
	fmt.Println("=====》[Client connection", strconv.Itoa(int(connection.GetConnectionID())), "stop]")
	fmt.Println(connection.GetProperty("web"))
	if err := connection.Send(0, []byte("{bye bye}")); err != nil {
		fmt.Println(err)
	}
}

func main() {
	Client := hnet.NewClient("0.0.0.0", 8999)
	Client.SetConnectionStartHook(ClientConnectionStartHook)
	Client.SetConnectionStopHook(ClientConnectionStopHook)
	//
	//Client.AddRouter(201, &ClientRouter1{})
	//Client.AddRouter(200, &ClientRouter{})
	Client.Start()

	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt, os.Kill)
	//sig := <-c
	//fmt.Println("===exit===", sig)
	//Client.Stop()
	//time.Sleep(time.Second * 20)
}
