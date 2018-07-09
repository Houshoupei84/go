package main

import "fmt"

func main(){

	a := make([]int , 5)
	printSlice("a ", a)

	//make的参数 第一个为slice的类型，第二个为len， 第三个为cap
	//// len(b)=0, cap(b)=5
	b := make([]int ,0, 5)
	printSlice("b ", b);

	c := b[ :2 ]

	printSlice(" c ", c)

	d := c[ 2 : 5]
	printSlice(" d ", d)
}


func printSlice(s string , x[]int){

	//这里注意换行的时间 标点符号要留在上一行 下面的这种写法会报错
	//fmt.Printf("%s len = %d cap = %d %v \n"
	//, s, len(x), cap(x), x)

	fmt.Printf("%s len = %d cap = %d %v \n", s, len(x), cap(x), x)
}
