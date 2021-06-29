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

/**
start a server
*/
func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name : %s, listenner at IP :%s, Port %d is starting\n", s.Name, s.IP, s.Port)
	go func() {

		//1 resolve a server address
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%s", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve server failed", err)
			return
		}
		//2 listener the address
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Printf("listener IPVersion:%s adress failed,%s", s.IPVersion, err)
			return
		}
		fmt.Printf("listener IPVersion:%s adress success", s.IPVersion)
		//3 Blocking  waiting  for client connection,processing(处理) the client  connection	of business(业务) (read and write)
		for {
			tcp, err := listener.AcceptTCP()
			if err != nil {
				fmt.Printf("Accept TCP err,%s", err)
				continue
			}
			//has established(建立)  a connection with the client ,
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := tcp.Read(buf)
					if err != nil {
						fmt.Printf("recive buf failed,%s", err)
						continue
					}

					//echo
					if _, err := tcp.Write(buf[:cnt]); err != nil {
						fmt.Printf("write bacl failed,%s", err)
						continue
					}
				}
			}()

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
		Port:      8888,
	}
}
