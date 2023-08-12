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

func CompareDaysOfWeek(number int) string {
	var day string

	switch {
	case number == 1:
		day = "Sunday"
		fallthrough // return next condition
	case number == 2:
		day = "Monday"
	default:
		day = "invalid"
	}

	return day
}

func main() {
	fmt.Println("Switch")
	fmt.Println(daysOfWeek(1))

	day := CompareDaysOfWeek(2)
	fmt.Println(day)
}
