package main

import "fmt"

func main() {
	tasks := make(chan int, 50)
	results := make(chan int, 50)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 50; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 50; i++ {
		result := <-results
		fmt.Println(result)
	}
}

func worker(tasks <-chan int, results chan<- int) {
	// receive(<-chan) and send(chan<-)
	for number := range tasks {
		results <- fibonacci(number)
	}

}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}
