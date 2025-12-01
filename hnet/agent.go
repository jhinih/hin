package hnet

import (
	"encoding/json"
	"errors"
	"github.com/jhinih/hin/hinterface"
	"sync"
)

type Agent struct {
	Identity map[string]interface{}

	ClientManager hinterface.IClientManager
	IdentityLock  sync.RWMutex
	Server        hinterface.IServer

	ServerManager hinterface.IServerManager
}
type BaseAgent struct {
	Identity map[string]interface{}

	ClientManager map[uint32]hinterface.IClient
	Lock          sync.RWMutex
	Server        hinterface.IServer

	ServerManager map[uint32]hinterface.IServer
}

func NewAgent() hinterface.IAgent {
	return &Agent{
		Identity:      make(map[string]interface{}),
		ClientManager: NewClientManager(),
		IdentityLock:  sync.RWMutex{},

		ServerManager: NewServerManager(),
	}
}
func (a *Agent) Start() {
	if a.Server != nil {
		a.Server.Start()
		a.ClientManager.StartClient()
	} else {
		a.ClientManager.StartClient()
		a.ServerManager.StartServer()
	}
}
func (a *Agent) Stop() {
	if a.Server != nil {
		a.Server.Stop()
	} else {
		a.ServerManager.ClearServer()
	}
}
func (a *Agent) Run() {
	a.Start()
}

func (a *Agent) AddServer(id uint32, server hinterface.IServer) {
	a.ServerManager.Add(server)
}
func (a *Agent) AddClient(id uint32, client hinterface.IClient) {
	a.ClientManager.Add(client)
}

func (a *Agent) SetIdentity(key string, value interface{}) {
	a.IdentityLock.Lock()
	defer a.IdentityLock.Unlock()

	a.Identity[key] = value
}
func (a *Agent) GetIdentity(key string) interface{} {
	a.IdentityLock.RLock()
	defer a.IdentityLock.RUnlock()

	if value, ok := a.Identity[key]; ok {
		return value
	} else {
		return errors.New("identity not exist")
	}
}
func (a *Agent) RemoveIdentity(key string) {
	a.IdentityLock.Lock()
	defer a.IdentityLock.Unlock()

	delete(a.Identity, key)
}

func (a *Agent) GetAllIdentity() map[string]interface{} {
	a.IdentityLock.RLock()
	defer a.IdentityLock.RUnlock()
	result := make(map[string]interface{})
	for k, v := range a.Identity {
		result[k] = v
	}
	return result
}

func (a *Agent) GetIdentityJSON() ([]byte, error) {
	a.IdentityLock.RLock()
	defer a.IdentityLock.RUnlock()

	return json.Marshal(a.Identity)
}
