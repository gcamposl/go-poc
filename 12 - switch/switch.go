package main

import "fmt"

func daysOfWeek(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	default:
		return "invalid"
	}
}

func main() {
	fmt.Println("Switch")
	fmt.Println(daysOfWeek(1))
}
