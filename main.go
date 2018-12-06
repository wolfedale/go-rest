package main

import (
	"flag"
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
	addr := flag.String("addr", ":5000", "HTTP network address")
	flag.Parse()

	srv := &http.Server{
		Addr:    *addr,
		Handler: routes(),
	}
	log.Fatal(srv.ListenAndServe())

	return nil
}
