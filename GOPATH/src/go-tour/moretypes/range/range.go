package main

import "fmt"

var pow = []int {1 , 2 , 4 , 8 , 16 , 32 , 64 , 128 }


func main(){

	//i 为 slice的下标
	//v 为 下标对应的元素

	for i, v := range pow {
		fmt.Printf("2 ** %d = %d\n", i, v)
	}
}
