package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Hello is a simple handler
type Hello struct {
	l *log.Logger
}

//NewHello creates a new hello handler with the given logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServerHTTP implements the go http.Handler interface
// golang net/http package
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//log.Println("Hello World")
	h.l.Println("Hello World")

	//read the body
	//d is data
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		//It has two method for Error

		//1. Method
		http.Error(rw, "Ooops", http.StatusBadRequest)

		//2. Method
		//rw.WriteHeader(http.StatusBadRequest)
		//rw.Write([]byte("Ooops"))
		return
	}

	//log.Printf("Data %s\n", d)
	// write the response
	fmt.Fprintf(rw, "Hello %s", d)
}
