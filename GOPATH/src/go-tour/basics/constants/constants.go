package main

import "fmt"

//const 可以定义int string bool
const Pi = 3.14

func main(){

	const World = "世界"

	fmt.Println("Hello", "World")
	fmt.Println("Happy", Pi, "Day")

	const Truth = true

	fmt.Println("Go rules ?",  Truth)
}
