package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexer(write("hello channel 1"), write("by channel 2"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexer(channel1, channel2 <-chan string) <-chan string {
	outChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-channel1:
				outChannel <- message
			case message := <-channel2:
				outChannel <- message
			}
		}
	}()

	return outChannel
}

func write(text string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- fmt.Sprintf("received value: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return channel
}
