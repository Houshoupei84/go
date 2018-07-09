package main

import (
	"net"
	"time"
	"fmt"
	"strings"
	"bufio"
	"log"
)

func echo (c net.Conn, shout string, delay time.Duration){

	fmt.Fprintln(c , "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t",strings.ToLower(shout))
}

//go
func handleConn(c net.Conn){

	//服务端优先从客户端读出数据 然后再返还回去
	input := bufio.NewScanner(c)

	for input.Scan(){
		 echo(c, input.Text(), 1* time.Second)
	}
	c.Close()
}

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
		go handleConn(conn)
	}
}
