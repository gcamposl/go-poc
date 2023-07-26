package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main()	{
	fmt.Println("writing from file main")
	auxiliar.Write()

	error := checkmail.ValidateFormat("x@gmail.com")

	fmt.Println(error)
}