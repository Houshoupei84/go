package main

import "fmt"

func main(){

	var a []int

	printSlice("a" , a);

	a =  append(a, 0);
	printSlice("a" , a)

	//append向 slice中添加一个元素
	//每添加一个元素 len 和 cap都加1
	a = append(a , 1)
	printSlice("a",  a)

	//这种一下加多个的 cap 比len多了1 ！！！！！！
	a = append(a , 2, 3, 4 )
	printSlice("a", a)
}

func printSlice(s string, x []int) {
	//打印slice的值的格式为  %v
	fmt.Printf("slice is %s lenth%d cap %d  slice value %v \n" , s, len(x), cap(x), x)
}
