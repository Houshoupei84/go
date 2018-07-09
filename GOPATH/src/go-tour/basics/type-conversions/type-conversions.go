package main

import (
	"math"
	"fmt"
)

func main(){

	var x, y int = 3, 4


	var f float64 = math.Sqrt( float64(x*x + y*y) )

	//打印浮点型
	fmt.Printf("%f\n" ,  f)

	var z  int = int ( f )

	//打印整形 注意f 默认被按照整形输出
	fmt.Println(x, y, z, f)
}
