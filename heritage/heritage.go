package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      uint
	height   uint
}

type student struct {
	person  // heritage
	course  string
	college string
}

func main() {
	fmt.Println("Heritage")

	person1 := person{"gabriel", "lopes", 26, 171}
	fmt.Println(person1)

	student1 := student{person1, "bcc", "unoeste"}
	fmt.Println(student1)
	fmt.Println(student1.name)
	fmt.Println(student1.person.name)
}
