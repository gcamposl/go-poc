package main

import "fmt"

func signReverse(number int) int {
	return number * -1
}

func signWithPointerReverse(number *int) {
	*number = *number * -1
}

func main() {
	fmt.Println("functions with pointers")
	number := 7
	reverseNumber := signReverse(number)
	fmt.Println(number)
	fmt.Println(reverseNumber)

	newNumber := 8
	fmt.Println(newNumber)
	signWithPointerReverse(&newNumber)
	fmt.Println(newNumber)
}
