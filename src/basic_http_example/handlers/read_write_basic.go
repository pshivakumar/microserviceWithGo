package handler

import (
	"fmt"
	"net/http"
)

//Hello - struct
type Hello struct{}

//NewHello - func
func NewHello() http.Handler {
	return Hello{}
}

//Hello - interface implemented
func (h Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello World")
}
