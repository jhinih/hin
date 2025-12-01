package hinterface

type IClientManager interface {
	Add(IClient)
	Remove(IClient)
	Get(uint64) (IClient, error)
	Len() int

	ClearClient()
	StartClient()
}
