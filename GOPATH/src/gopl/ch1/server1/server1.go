package main

import (
	"net/http"
	"log"
	"fmt"
)

func main(){

	http.HandleFunc("/", handler)
	//http.HandleFunc("/hello", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil));
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "URL.Path %q \n", r.URL.Path);
}