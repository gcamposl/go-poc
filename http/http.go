package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home page"))
}

func main() {
	http.HandleFunc("/home", home)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("users page"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
