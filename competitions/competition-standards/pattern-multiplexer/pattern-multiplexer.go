package main

import (
	"fmt"
	"time"
)

func write(text string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- fmt.Sprintf("received value: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return channel
}

func main() {
	test := write("hello")
	fmt.Println(test)
}
