package main

import "fmt"

func add (number1 int8, number2 int8) int8 {
	return number1 + number2
}

// double returns
func calculate (number3, number4 int8) (int8, int8) {
	sum := number3 + number4
	subtraction := number3 - number4
	return sum, subtraction
}

func main () {
	soma := add(2, 2)
	fmt.Println(soma)

	var f1 = func (text string) string {
		fmt.Println(text)
		return "return func -> f1"
	}

	result := f1("execution function -> f1")
	fmt.Println(result)
}
