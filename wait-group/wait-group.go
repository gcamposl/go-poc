package main

import (
	"fmt"
	"sync"
	"time"
)

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func main() {
	// sync goroutines
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		write("func 1")
		waitGroup.Done() // waitGroup--
	}()

	go func() {
		write("func 2")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
