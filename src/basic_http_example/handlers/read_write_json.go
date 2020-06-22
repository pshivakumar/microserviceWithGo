package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Greeting - struct
type Greeting struct {
	Message string
}

//NewGreeting - func
func NewGreeting(m string) *Greeting {
	return &Greeting{m}
}

//GenericHello - func
func GenericHello(rw http.ResponseWriter, r *http.Request) {
	var g Greeting

	decoder := json.NewDecoder(r.Body)
	decodeErr := decoder.Decode(&g)

	if decodeErr != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(rw)
	encodeErr := encoder.Encode(g)

	if encodeErr != nil {
		fmt.Println("Error....")
		return
	}
}
