package main

import (
	"net"
	"log"
	"io"
	"time"
	"fmt"
)

func main(){
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn , err := listener.Accept()
		fmt.Println("Accept conn", conn)
		if err != nil {
			log.Print(err)
			continue
		}
		//go 因为没有go 所以第一个链接来了之后，就在handleConn的死循环里处理。不再能新增加Accept
		handleConn(conn)
	}
}


func handleConn(c 	net.Conn){

	defer c.Close()

	for{
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))

		fmt.Println("handleConn conn WriteString", c)
		if err != nil{
			return
		}

		time.Sleep(1* time.Second)
	}
}
