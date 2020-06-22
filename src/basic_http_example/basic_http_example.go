package main

import (
	"net/http"

	handler "./handlers"
)

func main() {
	addr := "127.0.0.1:9090"

	helloHandler := handler.NewValidate(handler.NewHello())
	genericHelloHandler := handler.NewValidate(handler.NewGreeting())

	http.Handle("/hello", helloHandler)
	http.Handle("/generichello", genericHelloHandler)

	http.ListenAndServe(addr, nil)
}
