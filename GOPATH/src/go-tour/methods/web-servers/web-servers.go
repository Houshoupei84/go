package main

import (
	"net/http"
	"fmt"
	"log"
)

type Hello struct {

}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request){

		//这里注意 w指向的是 发起请求的http 所以hello输出在浏览器里
		fmt.Fprint(w, "Hello!")
		//fmt.Printf("Hello!!!!\n")
}

func main(){

	var h Hello

	err := http.ListenAndServe("localhost:4000", h)

	if err != nil {
		log.Fatal(err)
	}
}