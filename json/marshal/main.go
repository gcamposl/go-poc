package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Age  uint   `json:"age"`
}

func main() {
	d := dog{"xpto", "york", 2}
	dogInJSON, err := json.Marshal(d)

	if err != nil {
		log.Fatalf("error: %f", err)
	}

	fmt.Println(dogInJSON)
	fmt.Println(bytes.NewBuffer(dogInJSON))

	d2 := map[string]string{
		"name": "rex",
		"race": "pitbul",
	}

	dog2InJSON, err := json.Marshal(d2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dog2InJSON)
	fmt.Println(bytes.NewBuffer(dog2InJSON))
}
