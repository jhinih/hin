package hnet

import "github.com/jhinih/hin/hinterface"

type BaseRouter struct {
}

func NewBaseRouter() *BaseRouter {
	return &BaseRouter{}
}

func (b *BaseRouter) PreHandle(hinterface.IRequest) {
}
func (b *BaseRouter) Handle(hinterface.IRequest) {
}
func (b *BaseRouter) PostHandle(hinterface.IRequest) {
}
