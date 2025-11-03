package hnet

import (
	"errors"
	"fmt"
	"github.com/jhinih/hin/hinterface"
	"strconv"
	"sync"
)

type ConnectionManager struct {
	Connections map[uint32]hinterface.IConnection
	Lock        sync.RWMutex
}

func NewServerConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		Connections: make(map[uint32]hinterface.IConnection),
	}
}
func NewClientConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		Connections: make(map[uint32]hinterface.IConnection),
	}
}

func (c *ConnectionManager) Add(Connection hinterface.IConnection) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	c.Connections[Connection.GetConnectionID()] = Connection
	fmt.Println("[connection manager add connection", Connection.GetConnectionID(), "]")

}
func (c *ConnectionManager) Remove(Connection hinterface.IConnection) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	delete(c.Connections, Connection.GetConnectionID())
	fmt.Println("[connection manager delete connection:", Connection.GetConnectionID(), "]")

}
func (c *ConnectionManager) Get(connectionID uint32) (hinterface.IConnection, error) {
	c.Lock.RLock()
	defer c.Lock.RUnlock()

	if val, ok := c.Connections[connectionID]; ok {
		return val, nil
	} else {
		return nil, errors.New("connection " + strconv.Itoa(int(connectionID)) + "is not found")
	}
}
func (c *ConnectionManager) Len() int {
	return len(c.Connections)
}
func (c *ConnectionManager) ClearConnection() {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	for connectionID, connection := range c.Connections {
		connection.Stop()
		delete(c.Connections, connectionID)
	}
	fmt.Println("[clear all connection]")

}
