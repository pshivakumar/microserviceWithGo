package main

import (
	"net/http"

	handler "./handlers"
)

func main() {
	addr := "127.0.0.1:9090"
	http.HandleFunc("/helloworld", handler.Hello)
	http.HandleFunc("/generichello", handler.GenericHello)

	http.ListenAndServe(addr, nil)
}
