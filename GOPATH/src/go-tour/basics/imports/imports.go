package main

import (
	"fmt"
	"math"
)

func main(){

	//和y做比较，如果y比较大，那么返回最接近x并且是大于x的值。
	fmt.Printf("Now you have %g problems \n", math.Nextafter(2,3))
	//和y做比较，如果y比较小，那么返回最接近x并且是小于x的值。
	fmt.Printf("Now you have %g problems \n", math.Nextafter(2,1))

}
