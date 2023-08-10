package main

import "fmt"

func main() {
	func(text string) {
		fmt.Println(text)
	}("anonymous functions")

	result := func(text string) string {
		return fmt.Sprintf("return -> %s", text)
	}("anonymous functions")

	fmt.Println(result)
}
