package main

import "fmt"

func function1() {
	fmt.Println("execution 1")
}

func function2() {
	fmt.Println("execution 2")
}

func isAproved(n1, n2 float32) bool {
	defer fmt.Println("calcuted average. wait for the result")
	if average := (n1 + n2) / 2; average >= 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println("defer")
	defer function1()
	function2()

	result := isAproved(5, 7)
	if result {
		fmt.Println("aproved!")
	} else {
		fmt.Println("reproved!")
	}
}
