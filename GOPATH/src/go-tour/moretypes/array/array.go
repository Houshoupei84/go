package main

import "fmt"

func main(){

	var a[2] string

	a[0] = "Hello"
	a[1] = "World"

	//这样只打印 每个元素的值
	fmt.Println(a[0], a[1])
	//这里把数组打印出来 包括数组的值
	fmt.Println(a)
}
