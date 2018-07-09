package main

import "fmt"

func sum( a []int , c chan int){

	sum := 0

	for _ , v := range a {
		sum += v
	}

	//但是可以在每个协程里面调用管道
	c <- sum //将和送入c
}

func main(){

	a := []int { 7, 2, 8, -9, 4, 0}

	//注意创建管道的方法 没有指定缓存的个数 所以不能在主线程中往管道里写值
	c := make(chan int)

	//c <- 1

	//17
	go sum(a[ : len(a)/2], c)

	//-5
	go sum(a[len(a)/2 : ], c)

	//注意x为 -5  y为17 说明管道是后入先出的栈类型
	x, y := <- c, <- c //从 c中获取值

	////-5 17 12
	//x := <- c
	//y := <- c

	fmt.Println(x, y, x+y)
}

