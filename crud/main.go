package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.CreateUser).Methods(http.MethodGet)

	fmt.Println("Listening on port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
}
