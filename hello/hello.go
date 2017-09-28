package hello

import (
	"fmt"
	"log"
	"net/http"

	"github.com/leonegeo/demoapp-gocommon"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	log.Printf("hello got /hello")
	fmt.Fprintf(w, "Hello.")
}

func handleHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi.")
}

func Run() error {

	port, err := common.GetPort()
	if err != nil {
		return err
	}

	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/hi", handleHi)

	return http.ListenAndServe(":"+port, nil)
}
