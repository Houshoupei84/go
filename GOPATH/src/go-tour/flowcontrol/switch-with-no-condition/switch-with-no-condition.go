package main

import (
	"time"
	"fmt"
)

func main(){

	t := time.Now()

	switch {

	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good evening")
	}
}