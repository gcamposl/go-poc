package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var value1 int = 1
	var value2 int = value1
	fmt.Println(value1, value2)

	value1++
	fmt.Println(value1, value2)

	var value3 int = 1
	var value4 *int = &value3

	fmt.Println(value3, value4)
	value3++
	fmt.Println(value3, *value4)

	*value4 = 100
	fmt.Println(value3, *value4)
}
