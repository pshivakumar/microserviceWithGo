package handler

import (
	"net/http"
)

// Validate - struct
type Validate struct {
	next http.Handler
}

// NewValidate - func
func NewValidate(next http.Handler) http.Handler {
	return Validate{next: next}
}

func (v Validate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	v.next.ServeHTTP(rw, r)
}