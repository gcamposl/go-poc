package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and Slices")

	var array [5]string
	fmt.Println(array)

	array[0] = "zero"
	fmt.Println(array)
	fmt.Println(array[0])

	array2 := [5]int{0, 1, 2, 3, 4}
	fmt.Println(array2)

	array3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(array3)

	// slices

	slice := []int{1, 2, 3}
	slice[2] = 5
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 6)
	fmt.Println("new slice:", slice)

	slice = array3[1:4]
	fmt.Println(slice)

	// intern array

	slice2 := make([]int, 10, 11)
	fmt.Println(slice2)

	slice2 = append(slice2, 1)
	slice2 = append(slice2, 2)

	fmt.Println(len(slice2)) // lenght
	fmt.Println(cap(slice2)) // capacity

	slice3 := make([]int, 10)
	fmt.Println(slice3)
	slice2 = append(slice3, 3)
	fmt.Println(len(slice3)) // lenght
	fmt.Println(cap(slice3)) // capacity}
}
