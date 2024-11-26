package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		w.Write([]byte("failed to converter user in struct - " + err.Error()))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("failed to connect to database - " + err.Error()))
		return
	}
	defer db.Close()

	statment, err := db.Prepare("insert into users (name, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("failed to create statement - " + err.Error()))
		return
	}
	defer statment.Close()

	insertion, err := statment.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("failed to execute statement - " + err.Error()))
		return
	}

	insertionId, err := insertion.LastInsertId()
	if err != nil {
		w.Write([]byte("failed to get the last inserted id - " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("inserted id: %d", insertionId)))
}

// get all users in database
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("failed to connect to database"))
	}
	defer db.Close()
	sql := "select * from users"
	lines, err := db.Query(sql)
	if err != nil {
		w.Write([]byte("failed to get all users"))
	}
	defer lines.Close()
	var users []user
	for lines.Next() {
		var user user
		if err := lines.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("failed to scan users"))
			return
		}
		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("failed to converter user to json"))
		return
	}
}

// get a specific user by id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["id"], 10, 32)
	if err != nil {
		w.Write([]byte("failed to converter parameter to int"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("failed to connect to database"))
		return
	}

	line, err := db.Query("select * from users where id = ?", ID)
	if err != nil {
		w.Write([]byte("failed to get user"))
		return
	}

	var user user
	if line.Next() {
		if err := line.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("failed to scan user"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("failed to converter user to json"))
		return
	}
}

// update data users in database
func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	id, err := strconv.ParseUint(parameters["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error when converter parameter"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Error when read request body"))
		return
	}

	var user user
	if err := json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Error when deserialize request body"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Error when try to connect to database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update users set name = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Error when try to create statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, id); err != nil {
		w.Write([]byte("Error when try to exec statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteUser remove user from database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	Id, err := strconv.ParseUint(parameters["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error when converter parameter."))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Error when try to connect to database."))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from users where id = ?")
	if err != nil {
		w.Write([]byte("Error when try to create statement."))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(Id); err != nil {
		w.Write([]byte("Error when try to exec statement."))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
