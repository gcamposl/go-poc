package main

import (
	"fmt"
	"time"
)

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}

func main() {
	fmt.Println("channels")
	channel := make(chan string)

	go func() {
		write("hello", channel)
		fmt.Println("write finished")
	}()

	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("end of program")
}
