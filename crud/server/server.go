package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// insert user into database
// test
func CreateUser(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("failed to read body request!"))
		return
	}

	var user user
	if err = json.Unmarshal(request, &user); err != nil {
		w.Write([]byte("failed to converter user in struct"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("failed to connect to database"))
		return
	}
	defer db.Close()

	statment, err := db.Prepare("insert into users (name, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("failed to create statement"))
		return
	}
	defer statment.Close()

	insertion, err := statment.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("failed to execute statement"))
		return
	}

	insertionId, err := insertion.LastInsertId()
	if err != nil {
		w.Write([]byte("failed to get the last inserted id"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("inserted id: %d", insertionId)))
}

// get all users in database
func getAllUsers(w http.ResponseWriter, r *http.Request) {

}

// get a specific user by id
func getUserById(w http.ResponseWriter, r *http.Request) {

}
