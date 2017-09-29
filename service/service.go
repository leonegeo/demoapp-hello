package service

import (
	"fmt"
	"net/http"
)

// HelloHandler implements "/hi" endpoint
type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//log.Printf("hello got /hello: %s", r.URL)
	fmt.Fprintf(w, "Hello.")
}

// HiHandler implements "/hi" endpoint
type HiHandler struct{}

func (h *HiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//log.Printf("hi got /hi")
	fmt.Fprintf(w, "Hi.")
}

// Handlers returns the handlers
func Handlers() http.Handler {

	helloHandler := &HelloHandler{}
	hiHandler := &HiHandler{}

	mux := http.NewServeMux()
	mux.Handle("/hello", helloHandler)
	mux.Handle("/hi", hiHandler)

	return mux
}
