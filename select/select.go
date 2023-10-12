package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("select")
	channel, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel <- "channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "channel 2"
		}
	}()

	for {
		select {
		case channelMessage := <-channel:
			fmt.Println(channelMessage)
		case channelMessage2 := <-channel2:
			fmt.Println(channelMessage2)
		}
	}
}
