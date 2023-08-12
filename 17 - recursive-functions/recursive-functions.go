package main

import "fmt"

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {
	fmt.Println("recursive functions")
	position := uint(20)
	for i := uint(0); i < position; i++ {
		fmt.Println(fibonacci(i))
	}
}
