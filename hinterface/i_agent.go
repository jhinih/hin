package hinterface

type IAgent interface {
	Start()
	Stop()
	Run()

	AddServer(uint32, IServer)
	AddClient(uint32, IClient)

	SetIdentity(key string, value interface{})
	GetIdentity(key string) interface{}
	RemoveIdentity(key string)
	GetAllIdentity() map[string]interface{}
	GetIdentityJSON() ([]byte, error)
}
