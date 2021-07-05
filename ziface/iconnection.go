package ziface

import "net"

type IConnection interface {
	//start the connection
	Start()
	//stop the connection
	Stop()
	//get  a TCP connection
	GetTCPConnection() *net.TCPConn
	//get the connection id
	GetConnId() uint32
	//get the  remote connection address
	RemoteAddr() net.Addr
	//send message to remote client
	SendMsg(data []byte) error
}

//delimit a func to hand business
type HandleFunc func(net.Conn, []byte, int) error
