package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":     "gabriel",
		"lastName": "lopes",
	}

	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"user": {
			"firstName": "gabriel",
			"lastname":  "lopes",
		},
		"course": {
			"name":   "Computer Science",
			"campus": "campus 1",
		},
	}
	fmt.Println(user2)
	delete(user2, "course")
	fmt.Println(user2)

	user2["document"] = map[string]string{
		"cpf":   "12345678900",
		"state": "sp",
	}
	fmt.Println(user2)
}
