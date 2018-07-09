package main

import "fmt"

//相同的数据类型为 int
var i , j int = 1, 2

func main(){

	//不同的数据类型 bool 和 string
	var c, python, java = true , false , "no !"

	fmt.Println(i , j, c, python, java)
}
