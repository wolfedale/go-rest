package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func routes() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/users", adminOnly(getUsers)).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users", addUser).Methods("POST")
	router.HandleFunc("/users", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	router.HandleFunc("/", adapt(getUser, logRequest(), secureHeaders()))

	return router
}
