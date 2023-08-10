package main

import "fmt"

// numbers ...int == slice
func calculate(numbers ...int) {
	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println(total)
}

func write(text string, numbers ...int) {
	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println(text, total)
}

func main() {
	calculate(1, 2, 3)
	calculate(1, 2, 3, 4)
	write("test", 1, 2, 3)
}
