package znet

import (
	"zinx_learning/ziface"
)

type BaseRouter struct{}

//define empty methods to override

func (base *BaseRouter) PreHandle(requeest ziface.IRequest) {}

func (base *BaseRouter) Handle(requeest ziface.IRequest) {}

func (base *BaseRouter) PostHandle(request ziface.IRequest) {}
