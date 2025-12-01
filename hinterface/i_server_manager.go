package hinterface

type IServerManager interface {
	Add(IServer)
	Remove(IServer)
	Get(uint64) (IServer, error)
	Len() int
	ClearServer()
	StartServer()
}
