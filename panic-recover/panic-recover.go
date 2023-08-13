package main

import "fmt"

func tryRecoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("attempt to recover execution")
	}
}

func isAproved(n1, n2 float32) bool {
	defer tryRecoverExecution()
	average := (n1 + n2) / 2
	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	panic("6 NOOO")
}

func main() {
	fmt.Println("panic and recover")
	fmt.Println(isAproved(6, 6))
	fmt.Println("pos execution")
}
