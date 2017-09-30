package main

import (
	"log"
	"net/http"

	common "github.com/leonegeo/demoapp-gocommon"
)

func main() {

	port, err := common.GetPort()
	if err != nil {
		log.Fatal(err)
	}

	//log.Fatal(http.ListenAndServe(port, server.Handlers()))

	s := &http.Server{
		Addr:    port,
		Handler: Handlers(),
	}

	log.Fatal(s.ListenAndServe())
}
