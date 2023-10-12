package main

import "fmt"

func main() {
	channel := make(chan string, 2) // buffer
	channel <- "hello"
	channel <- "hello again"

	message := <-channel  // receive value 1
	message2 := <-channel // '' 2

	fmt.Println(message)
	fmt.Println(message2)
}
