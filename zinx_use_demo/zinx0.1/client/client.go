package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("clint is starting")
	time.Sleep(1 * time.Second)
	//1.connect to server ,get a connection
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Printf("client connect  server failed ,%s", err)
		return
	}

	for {
		_, err := conn.Write([]byte("Hello Zinx V0.1.."))
		if err != nil {
			fmt.Printf("write conn  failed,%s\n", err)
		}
		buf := make([]byte, 512)
		read, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("read conn  failed,%s\n", err)
		}
		fmt.Printf("server call back: %s, cnt = %d\n", buf, read)
		time.Sleep(1 * time.Second)

	}

}
