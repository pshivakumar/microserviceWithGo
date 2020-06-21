package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

	umErr := json.Unmarshal(body, &g)

	if umErr != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(g)

	if err != nil {
		fmt.Println("Error....")
	}
	fmt.Fprint(rw, string(data))
}
