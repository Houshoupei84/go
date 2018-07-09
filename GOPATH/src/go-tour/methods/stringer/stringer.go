package main

import "fmt"

type Person struct {

	Name string
	Age  int
}

func (p Person) String() string {

	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main(){

	a := Person{"qq", 19}
	z := Person{"xx", 19}
	fmt.Println(a, z)
}