package main

import (
	"time"
	"fmt"
)

//定义结构体 MyError
type MyError struct {
	When time.Time
	What string
}

//MyError 定义 Error接口
func (e* MyError) Error() string{

	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

/*
返回值error 很奇怪  为啥&MyError 会变成返回值error？？？？？？？？
*/
func run() error {

	return &MyError{
		time.Now,
		"it didnot work",
	}
}

func main(){

	if err := run(); err != nil{
		fmt.Println(err)
	}
}