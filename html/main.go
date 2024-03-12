package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type user struct {
	Name  string
	Email string
}

var templates *template.Template

func main() {
	// get all files with .html extensions
	templates = template.Must(template.ParseGlob("*.html"))

	u := user{
		"xpto",
		"xpto@xpto.xpto",
	}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("listening on 5000 port...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
