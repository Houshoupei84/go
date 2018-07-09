package main

import (
	"math"
	"fmt"
)

//Abser 为一个名字为Abs 返回值为float64的接口
type Abser interface{
	Abs() float64
}

func main(){

	var a Abser

	f := MyFloat(-math.Sqrt2)

	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())

	a = &v
	fmt.Println(a.Abs())

	//a = v
	//fmt.Println(a.Abs())
}

type MyFloat float64

//对象上调用Abs
func(f MyFloat) Abs() float64{
	if f < 0{
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X , Y float64
}


//指针上 调用Abs
func (v* Vertex) Abs() float64{

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
