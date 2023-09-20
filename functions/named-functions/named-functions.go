package main

import "fmt"

func calculate(number1, number2 int) (sum int, subtraction int) {
	sum = number1 + number2
	subtraction = number1 - number2
	return
}

func main() {
	sum, subtraction := calculate(2, 1)
	fmt.Println(sum, subtraction)
}
