package main

import (
	"time"
	"fmt"
)

func main(){

	go spinner(100 * time.Microsecond)

	const n = 45

	fibN := fib(n)

	fmt.Printf("\r  Fibonacci(%d) = %d\n", n ,fibN)
}


func spinner (delay time.Duration){

	//注意两层循环，第一层是个死循环 除非main  goroutine 结束
	for{
		//
		for _, r := range `-\|/` {
			fmt.Printf("\r %c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int{

	if(x < 2){
		return x
	}
	return fib(x-1) + fib(x -2)
}

