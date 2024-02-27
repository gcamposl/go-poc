package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"-"` // - ignore tag
	Race string `json:"race"`
	Age  uint   `json:"age"`
}

func main() {
	dogInJSON := `{"name": "bento", "raca": "york", "age": 1}`

	var d dog

	if err := json.Unmarshal([]byte(dogInJSON), &d); err != nil {
		log.Fatal(err)
	}
	fmt.Println(dogInJSON)

	dog2InJSON := `{"name": "bentina", "raca": "border"}`

	d2 := make(map[string]string)

	if err := json.Unmarshal([]byte(dog2InJSON), &d2); err != nil {
		fmt.Println(d2)
	}
}
