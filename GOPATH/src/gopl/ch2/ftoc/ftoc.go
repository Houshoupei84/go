package main

import "fmt"

func main(){

	const freezIngF, boilingF = 32.0, 212.0

	fmt.Printf("%g f = %g c \n", freezIngF, fToc(freezIngF))
	fmt.Printf("%g f = %g c \n",boilingF, fToc(boilingF))
}

func fToc(f float64) float64 {

	return (f - 32) * 5 / 9
}
