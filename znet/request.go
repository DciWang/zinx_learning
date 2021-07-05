package znet

import "zinx_learning/ziface"

type Request struct {
	//a link has been established with the client
	Conn ziface.IConnection
	//the data for client request
	data []byte
}

func (r *Request) GetConnect() ziface.IConnection {
	return r.Conn
}

func (r *Request) GetData() []byte {
	return r.data
}
