package main

import "fmt"

func setValue(value int) {
	value = 0
}

func setPointer(pointer *int) {
	*pointer = 0
}

func main() {
	fmt.Println("Pointers")

	i := 1
	fmt.Println("initial:", i)

	setValue(i)
	fmt.Println("setValue:", i)

	setPointer(&i)
	fmt.Println("setPointer:", i)

	fmt.Println("pointer:", &i)
}
