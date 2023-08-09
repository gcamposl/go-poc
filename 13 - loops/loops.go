package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	for i := 0; i < 5; i++ {
		fmt.Println("Index:", i)
		//time.Sleep(time.Second)
	}

	j := 5
	for j < 11 { // like while
		fmt.Println("Index:", j)
		j++
		//time.Sleep(time.Second)
	}

	names := [3]string{"abc", "cba", "bca"}
	for index, name := range names {
		fmt.Println(index, ":", name)
	}

	for _, name := range names { // without index
		fmt.Println(name)
	}

	for _, letter := range "letters" { // without index
		fmt.Println(string(letter))
	}

	user := map[string]string{
		"name":     "test",
		"lastName": "escescescesc",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}

	for {
		fmt.Println("infinity loop")
	}
}
