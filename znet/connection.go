package znet

import (
	"fmt"
	"net"
	"zinx_learning/ziface"
)

type connection struct {
	//current connection
	Conn *net.TCPConn
	//current connection  id
	ConnID uint32
	//current connection status
	isClosed bool

	//current connection method to handle business
	// ziface.HandleFunc
	//delimit a channel Notify that the current connection has been exited
	ExitChan chan bool
	//current connection bind router to  handle business
	Router ziface.IRouter
}

//business method of read
func (c *connection) StartRead() {
	fmt.Printf("reader goroutine is running...connId:%d\n", c.ConnID)
	defer c.Stop()
	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Printf("connection read failed,%s\n", err)
			continue
		}

		/*	if err := c.handleAPI(c.Conn, buf, read); err != nil {
			fmt.Printf("connection handle  failed,%s,connection id :%s\n", err, c.ConnID)
			break
		}*/
		//get current connection's request
		request := Request{
			c,
			buf,
		}
		//the method of execute  business
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&request)
		//get be bind  router  correspond(对应的) connection
	}

}

func (c *connection) Start() {
	fmt.Printf("connection start()...connId:%d\n", c.ConnID)
	//start current connection to handle business of read data
	go c.StartRead()
	//todo  start current connection to handle business of write data
}

func (c *connection) Stop() {
	fmt.Printf("connection stop()...connId:%d\n", c.ConnID)
	//Determine whether it has been closed
	if c.isClosed {
		return
	}

	c.isClosed = true
	//close socket connection
	err := c.Conn.Close()
	if err != nil {
		fmt.Printf("close socket connect failed,err:%s\n", err)
	}
	//recycling resources
	close(c.ExitChan)
}

func (c *connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *connection) GetConnId() uint32 {
	return c.ConnID
}

func (c *connection) RemoteAddr() net.Addr {
	return nil
}

func (c *connection) SendMsg(data []byte) error {
	return nil
}

func NewConnection(conn *net.TCPConn, connId uint32, router ziface.IRouter) *connection {
	return &connection{
		Conn:     conn,
		ConnID:   connId,
		isClosed: false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
}
