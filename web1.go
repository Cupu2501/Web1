package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handlerHelloWorld)
	http.ListenAndServe(":8080", nil)
}

func handlerHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World !!!")
}