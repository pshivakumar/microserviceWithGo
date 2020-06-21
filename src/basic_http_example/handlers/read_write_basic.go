package handler

import (
	"fmt"
	"net/http"
)

//Hello - struct
func Hello(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello World")
}
