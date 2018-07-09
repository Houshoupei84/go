package main

import "fmt"

func main(){

	//第二个参数为缓冲长度
	//第一个参数为 chan类型
	c := make(chan int , 2)//, 2

	c <- 1
	c <- 2

	//在同一个线程里 管道是先入先出
	fmt.Println( <- c)
	fmt.Println( <- c)
}
