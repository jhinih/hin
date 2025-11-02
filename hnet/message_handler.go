package hnet

import (
	"fmt"
	"github.com/jhinih/hin/hinterface"
	"strconv"
)

type MessageHandler struct {
	APIs map[uint32]hinterface.IRouter

	WorkPollSize uint32
	TaskChan     []chan hinterface.IRequest
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		APIs:         make(map[uint32]hinterface.IRouter),
		WorkPollSize: 10,
		TaskChan:     make([]chan hinterface.IRequest, 1024),
	}
}

func (m *MessageHandler) DoMessageHandler(request hinterface.IRequest) {
	handler, ok := m.APIs[request.GetMsgID()]
	if !ok {
		panic("API " + strconv.Itoa(int(request.GetMsgID())) + " is no found")
	}
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (m *MessageHandler) AddRouter(msgID uint32, router hinterface.IRouter) {
	if _, ok := m.APIs[msgID]; ok {
		panic("repeat api " + strconv.Itoa(int(msgID)))
	}
	m.APIs[msgID] = router
	fmt.Println("[add api", strconv.Itoa(int(msgID))+"]")
}

func (m *MessageHandler) StartWorkPoll() {
	for i := 0; i < int(m.WorkPollSize); i++ {
		m.TaskChan[i] = make(chan hinterface.IRequest, 1024)
		go m.StartWorker(i, m.TaskChan[i])
	}
}
func (m *MessageHandler) StartWorker(workerID int, taskChan chan hinterface.IRequest) {
	fmt.Println("[worker", strconv.Itoa(workerID), "start]")
	for {
		select {
		case request := <-taskChan:
			m.DoMessageHandler(request)

		}
	}
}

func (m *MessageHandler) SendMsg2TaskQueue(request hinterface.IRequest) {
	workerID := request.GetConnection().GetConnectionID() % m.WorkPollSize
	fmt.Println("[connection", request.GetConnection().GetConnectionID(), "MsgID", request.GetMsgID(), "add to", workerID, "TaskQueue]")

	m.TaskChan[workerID] <- request
}
