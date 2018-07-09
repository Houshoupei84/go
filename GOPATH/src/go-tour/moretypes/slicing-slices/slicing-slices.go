package main

import "fmt"

func main(){

	p := []int{ 2, 3, 5, 7, 11, 13}
	fmt.Println("p == ", p)

	//输出 3 5 7  说明下表从0开始，上限是小于4而不是小于等于
	fmt.Println("p[1 :4] ==", p[ 1 : 4])

	//下标0 ---3 但是小于3  2 3 5
	fmt.Println("p[ :3]" , p[:3])

	//这里注意 4到最后 但是包含最后一个  因此输出 11  13
	fmt.Println("p[4 :]", p[ 4:])
}
