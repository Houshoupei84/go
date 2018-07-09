package main

import "fmt"

func main(){

	fmt.Println("counting")

	//输出结果是 9--0  说明是栈类型的 先进后出
	for i:= 0 ; i < 10 ; i++  {

		defer fmt.Println(i)
	}

	fmt.Println("done")
}
