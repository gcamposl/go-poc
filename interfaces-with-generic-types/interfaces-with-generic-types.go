package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("interfaces with generic types")

	generic(1)
	generic("text")
	generic(true)
	generic(1.7)

	map1 := map[interface{}]interface{}{
		float64(1): true,
		"string":   1,
		true:       int(1),
	}

	fmt.Println(map1)
	fmt.Println(1, "text", true, 3.2)
}
