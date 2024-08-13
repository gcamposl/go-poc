package server

import (
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

	fmt.Println(user)
}
