package main

import (
	"net"
	"log"
	"os"
	"io"
)

func main(){

	conn, err := net.Dial("tcp", "localhost:8000")

	if err != nil	{
		log.Fatal(err)
	}

	defer conn.Close()

	go mustCopy(os.Stdout, conn)

	//从os.Stdin 读入数据发送给 conn
	//服务端收到后
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}