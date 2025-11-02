package hinterface

type IConnectionManager interface {
	Add(IConnection)
	Remove(IConnection)
	Get(uint32) (IConnection, error)
	Len() int
	ClearConnection()
}
