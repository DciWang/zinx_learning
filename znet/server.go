package znet

import (
	"fmt"
	"net"
	"zinx_learning/ziface"
)

/**
define a server ,implements IServer interface
*/
type Server struct {
	//the server's name
	Name string
	//the server's ipVersion
	IPVersion string
	//the server's Ip
	IP string
	//the server's port
	Port int
}

func handleFun(con net.Conn, buf []byte, cunt int) error {
	//business echo
	fmt.Println("[Conn Handle CallBack to client]")
	_, err := con.Write(buf[:cunt])
	if err != nil {
		fmt.Printf("connection handle callcack failed")
		return err
	}
	return nil
}

/**
start a server
*/
func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name : %s, listenner at IP :%s, Port :%d is starting\n", s.Name, s.IP, s.Port)
	go func() {
		//1 resolve a server address
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%v", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve server failed", err)
			return
		}
		//2 listener the address
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Printf("listener IPVersion:%s adress failed,%s\n", s.IPVersion, err)
			return
		}
		fmt.Printf("listener IPVersion:%s adress success\n", s.IPVersion)
		//3 Blocking  waiting  for client connection,processing(处理) the client  connection	of business(业务) (read and write)
		var connId uint32
		connId = 0
		for {
			tcp, err := listener.AcceptTCP()
			if err != nil {
				fmt.Printf("Accept TCP err,%s", err)
				continue
			}
			//has established(建立)  a connection with the client ,
			/*		go func() {
					for {
						buf := make([]byte, 512)
						read, err := tcp.Read(buf)
						if err != nil {
							fmt.Printf("recive buf failed,%s\n", err)
							continue
						}
						fmt.Printf("server call back: %s, cnt = %d\n", buf, read)
						//echo
						if _, err := tcp.Write(buf[:read]); err != nil {
							fmt.Printf("write bacl failed,%s\n", err)
							continue
						}
					}
				}()*/
			//bind our business and connection
			newConnection := NewConnection(tcp, connId, handleFun)
			connId++
			//start connection
			newConnection.Start()
		}
	}()

}
func (s *Server) Stop() {
	//todo  stop or reclaim   resources and status

}
func (s *Server) Server() {

	s.Start()
	//todo  Do some business after starting the server
	//block
	select {}
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
}
