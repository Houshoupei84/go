package main

import "fmt"

func main(){

	pow := make([]int, 10)

	//这里只要索引
	for i := range pow  {
		pow[i] = 1 << uint(i)
	}

	//这里只要value _只是一个占位符不能输出
	for _, value := range pow {
		fmt.Printf("%d \n",  value)
	}
}
