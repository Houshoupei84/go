package main


import "fmt"

type Vertes struct{

	X int
	Y int
}


func main(){

	v := Vertes{1,2}

	v.X = 4

	fmt.Println(v.X)
}