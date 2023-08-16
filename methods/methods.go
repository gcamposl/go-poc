package main

import "fmt"

type person struct {
	name string
	age  uint8
}

func (p person) save() {
	fmt.Printf("saving person: %s", p.name)
}

func main() {
	person := person{"gabriel", 26}
	fmt.Println(person)
	person.save()
}
