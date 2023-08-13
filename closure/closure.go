package main

import "fmt"

func closure() func() {
	text := "closure"
	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	text := "main"
	fmt.Println(text)

	newFunction := closure()
	newFunction()
}
