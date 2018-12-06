package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// for testing purposes
	err := realMain()
	if err != nil {
		os.Exit(2)
	}
}

func realMain() error {
	srv := &http.Server{
		Handler: routes(),
	}
	log.Fatal(srv.ListenAndServe())

	return nil
}
