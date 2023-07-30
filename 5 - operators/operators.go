package main

import "fmt"

func main () {
	// ARITHMETIC
	// + - / * %
	sum := 1 + 1
	subtraction := 10 - 1
	division := 14 / 2
	multiplication := 10 / 5
	rest := 10 % 2

	fmt.Println(sum, subtraction, division, multiplication, rest)

	var number1 int16 = 1
	var number2 int16 = 2
	var sum2 int16 = number1 + number2
	fmt.Println(sum2)

	// ASSIGNMENT (ATRIBUIÇÃO)
	var text string = "explict"
	text2 := "explicit"
	fmt.Println(text, text2)

	// RELATIONAL OPERATORS
	fmt.Println(number1 < number2)
	fmt.Println(number1 > number2)
	fmt.Println(number1 == number2)		
	fmt.Println(number1 <= number2)		
	fmt.Println(number1 >= number2)		
	fmt.Println(number1 != number2)

	// LOGIC OPERATORS
	tr, fal := true, false
	fmt.Println(tr && fal) // and
	fmt.Println(tr || fal) // or
	fmt.Println(!tr) // denial

	// UNARY OPERATORS
	number := 1
	number++
	number--
	number += 1
	number -= 1
	number /= 1
	number *= 1
	number %= 1
	fmt.Println(number)

	// ternary operators DOESN'T EXIST
	if number > 1 {
		text = "greater than 1"
	} else {
		text = "less than 1"
	}
	
	fmt.Println(text)
}