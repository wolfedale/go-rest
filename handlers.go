package main

import (
	"log"
	"net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("getUsers")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	log.Println("getUser")
}

func addUser(w http.ResponseWriter, r *http.Request) {
	log.Println("addUser")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("updateUser")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("deleteUser")
}
