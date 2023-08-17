package main

import "fmt"

type person struct {
	name string
	age  uint8
}

func (p person) save() {
	fmt.Printf("saving person: %s\n", p.name)
}

func (p person) ofLegalAge(age uint8) bool {
	return p.age > 18
}

func main() {
	person1 := person{"gabriel", 26}
	fmt.Println(person1)
	person1.save()

	person2 := person{"campos", 13}
	person2.save()

	isLegalAge := person2.ofLegalAge(person2.age)
	fmt.Println(isLegalAge)
}
