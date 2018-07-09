package main


import "fmt"

type Vertes struct{

	X int
	Y int
}


func main(){

	v := Vertes{1,2}

	p := &v

	//这里注意 虽然是指针 这里用的仍然是. 而不是->
	p.X = 7

	fmt.Println(v)
}