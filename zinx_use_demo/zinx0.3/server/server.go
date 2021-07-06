package main

import (
	"fmt"
	"zinx_learning/ziface"
	"zinx_learning/znet"
)

type pingRouter struct {
	znet.BaseRouter
}

// test PerHandle
func (r *pingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("call router  peHandle...")
	_, err := request.GetConnect().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Printf("call bacl before err,err:%s\n", err)
	}

}

// test Handle
func (r *pingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call router  Handle...")
	_, err := request.GetConnect().GetTCPConnection().Write([]byte(" ping... ping... ping...\n"))
	if err != nil {
		fmt.Printf("ping  err,err:%s\n", err)
	}
}

// test PostHandle
func (r *pingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("call router  postHandle...")
	_, err := request.GetConnect().GetTCPConnection().Write([]byte("after  ping...\n"))
	if err != nil {
		fmt.Printf("call bacl after err,err:%s\n", err)
	}
}

func main() {
	s := znet.NewServer("[Zinx V0.0.3 ]")
	s.AddRouter(&pingRouter{})
	s.Server()
}
