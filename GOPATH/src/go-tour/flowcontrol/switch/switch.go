package main

import (
	"fmt"
	"runtime"
)


func main(){

	fmt.Println("Go Runs on ")

	switch  os := runtime.GOOS; os{

	case "darwin":
		fmt.Println("os x.")

	case "linux":
		fmt.Println("Linux")

	default:
		fmt.Printf("%s." , os)

	}
}