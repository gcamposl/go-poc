package main

import "fmt"

type user struct {
	name string
	age uint8
	address address
}

type address struct {
	street string
	number int
}

func main () {
	fmt.Println("structs")
	var usr user
	usr.name = "gabriel"
	usr.age = 26
	fmt.Println(usr.name, usr.age)

	ads := address{"street x", 7}
 	usr2 := user{"campos", 62, ads}
	fmt.Println(usr2)

	usr3 := user{name: "lopes"}
	fmt.Println(usr3)
	
}