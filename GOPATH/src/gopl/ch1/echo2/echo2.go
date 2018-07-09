package main

import (
	"os"
	"fmt"
)

func main(){

	//s , sep := "", ""

	for index, arg := range os.Args[1:]{


		//s += sep + arg
		fmt.Println(index)
		fmt.Println(arg)
		//sep = " "
	}
}


