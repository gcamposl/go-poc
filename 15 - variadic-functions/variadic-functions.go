package main

import "fmt"

func calculate(numbers ...int) {
	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println(total)
}

func main() {
	calculate(1, 2, 3)
	calculate(1, 2, 3, 4)
}
