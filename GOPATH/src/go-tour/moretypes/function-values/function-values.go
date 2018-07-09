package main

import (
	"math"
	"fmt"
)

func main(){

	//说明函数可以定义在另一个函数里面
	hypot := func(x, y float64) float64{

		return  math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
}
