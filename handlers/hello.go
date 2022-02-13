package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//log.Println("Hello World")
	h.l.Println("Hello World")
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

	fmt.Fprintf(rw, "Hello %s", d)
}
