package main

import (
	"math/cmplx"
	"fmt"
)

var (
	ToBe 	 bool 			= false
	Maxint   uint64 		= 1 << 64 -1
	z        complex128     = cmplx.Sqrt(-5 + 12i )
)

func main (){

	//%T 输出类型
	//%v 输出value
	const f = "%T ( %v ) \n"

	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, Maxint, Maxint)
	fmt.Printf(f, z, z)

}
