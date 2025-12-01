package hnet

import (
	"errors"
	"fmt"
	"github.com/jhinih/hin/hinterface"
	"strconv"
	"sync"
)

type IClientManager struct {
	Clients map[uint64]hinterface.IClient
	Lock    sync.RWMutex
}

func NewClientManager() *IClientManager {
	return &IClientManager{
		Clients: make(map[uint64]hinterface.IClient),
	}
}
func (c *IClientManager) Add(Client hinterface.IClient) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	c.Clients[Client.GetClientID()] = Client
	fmt.Println("[Client manager add Client", Client.GetClientID(), "]")

}
func (c *IClientManager) Remove(Client hinterface.IClient) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	delete(c.Clients, Client.GetClientID())
	fmt.Println("[Client manager delete Client:", Client.GetClientID(), "]")

}
func (c *IClientManager) Get(ClientID uint64) (hinterface.IClient, error) {
	c.Lock.RLock()
	defer c.Lock.RUnlock()

	if val, ok := c.Clients[ClientID]; ok {
		return val, nil
	} else {
		return nil, errors.New("Client " + strconv.Itoa(int(ClientID)) + "is not found")
	}
}
func (c *IClientManager) Len() int {
	return len(c.Clients)
}
func (c *IClientManager) ClearClient() {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	for ClientID, Client := range c.Clients {
		Client.Stop()
		delete(c.Clients, ClientID)
	}
	fmt.Println("[clear all Client]")

}
func (c *IClientManager) StartClient() {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	for _, Client := range c.Clients {
		Client.Start()
	}
	fmt.Println("[start all Client]")

}
