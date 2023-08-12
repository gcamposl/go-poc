package main

import "fmt"

func main() {
	fmt.Println("Control Structures")

	number := 5

	if number > 1 {
		fmt.Println("number > 1")
	} else {
		fmt.Println("number < 1")
	}

	if auxNumber := 1; auxNumber > 0 { // scopped
		fmt.Println("auxNumber > 0")
	}
}
