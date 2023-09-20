package main

import "fmt"

var n int

func init() {
	n = 1
	fmt.Println("execution into the init")
}

func main() {
	fmt.Println("execution into the main")
	fmt.Println(n)
}
