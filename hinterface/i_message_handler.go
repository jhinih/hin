package hinterface

// IMsgHandle Abstract layer of message management(消息管理抽象层)
type IMsgHandle interface {
	// Add specific handling logic for messages, msgID supports int and string types
	// (为消息添加具体的处理逻辑, msgID，支持整型，字符串)
	AddRouter(msgID uint32, router IRouter)
	AddRouterSlices(msgId uint32, handler ...RouterHandler) IRouterSlices
	Group(start, end uint32, Handlers ...RouterHandler) IGroupRouterSlices
	Use(Handlers ...RouterHandler) IRouterSlices

	StartWorkerPool()                    //  Start the worker pool
	SendMsgToTaskQueue(request IRequest) // Pass the message to the TaskQueue for processing by the worker(将消息交给TaskQueue,由worker进行处理)

	Execute(request IRequest) // Execute interceptor methods on the responsibility chain(执行责任链上的拦截器方法)
}
