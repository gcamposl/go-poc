package main

import (
	"errors"
	"fmt"
)

func main() {
	// INTEGERS NUMBERS
	var number int = -10 // int8, int64, int32, int64
	fmt.Println(number)

	var unumber uint = 10 // unsigned int - uint8, uint64, uint32, uint64
	fmt.Println(unumber)
	// alias
	var unumber32 rune = 32 // uint32 alias
	fmt.Println(unumber32)

	// byte == uint8
	var bnumber byte = 123
	fmt.Println(bnumber)

	// REAL NUMBERS
	var fnumber32 float32 = 3.2
	fmt.Println(fnumber32)

	var fnumber64 float64 = 6.4
	fmt.Println(fnumber64)

	// STRINGS
	// '' -> char (there is no char)
	// "" -> string
	var text string = "text text text"
	fmt.Println(text)

	char := 'a' // return ASC 2
	fmt.Println(char)

	// BOOLEAN
	var boolean bool = true // or false
	fmt.Println(boolean)

	// INITIALIZING VARIABLES
	var text2 string
	fmt.Println(text2) // " "

	var number2 int
	fmt.Println(number2) // 0

	// ERROR
	var erro error
	fmt.Println(erro) // nil == null (js, csharp, java...)

	var erro2 error = errors.New("internal error")
	fmt.Println(erro2)
}