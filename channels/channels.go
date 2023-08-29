package main

import (
	"fmt"
	"time"
)

func write(text string, channel chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("channels")
	channel := make(chan string)

	go write("hello", channel)

	fmt.Println("after")
	message := <-channel // receive channel
	fmt.Println(message)
}
