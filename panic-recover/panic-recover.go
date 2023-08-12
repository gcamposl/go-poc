package main

import "fmt"

func isAproved(n1, n2 float32) bool {
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
	fmt.Println(isAproved(6, 5))
	fmt.Println("pos execution")
}
