package hnet

import (
	"errors"
	"fmt"
	"github.com/jhinih/hin/hinterface"
	"strconv"
	"sync"
)

type IServerManager struct {
	Servers map[uint64]hinterface.IServer
	Lock    sync.RWMutex
}

func NewServerManager() *IServerManager {
	return &IServerManager{
		Servers: make(map[uint64]hinterface.IServer),
	}
}
func (c *IServerManager) Add(Server hinterface.IServer) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	c.Servers[Server.GetServerID()] = Server
	fmt.Println("[Server manager add Server", Server.GetServerID(), "]")

}
func (c *IServerManager) Remove(Server hinterface.IServer) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	delete(c.Servers, Server.GetServerID())
	fmt.Println("[Server manager delete Server:", Server.GetServerID(), "]")

}
func (c *IServerManager) Get(ServerID uint64) (hinterface.IServer, error) {
	c.Lock.RLock()
	defer c.Lock.RUnlock()

	if val, ok := c.Servers[ServerID]; ok {
		return val, nil
	} else {
		return nil, errors.New("Server " + strconv.Itoa(int(ServerID)) + "is not found")
	}
}
func (c *IServerManager) Len() int {
	return len(c.Servers)
}
func (c *IServerManager) ClearServer() {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	for ServerID, Server := range c.Servers {
		Server.Stop()
		delete(c.Servers, ServerID)
	}
	fmt.Println("[clear all Server]")

}
func (c *IServerManager) StartServer() {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	for _, Server := range c.Servers {
		Server.Start()
	}
	fmt.Println("[start all Server]")

}
