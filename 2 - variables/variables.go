package main

import "fmt"

func main() {
	var variable string = "explicit"
	variable2 := "implicit"
	fmt.Println(variable)
	fmt.Println(variable2)

	var (
		variable3 string = "var3"
		variable4 string = "var4"
		variable5 int = 0
	)
	fmt.Println(variable3, variable4, variable5)

	variable6, variable7, variable8 := "var6", "var7", 8
	fmt.Println(variable6, variable7, variable8)

	const const1 string = "constant 1"
	fmt.Println(const1)

	//invert values without aux variable
	variable6, variable7 = variable7, variable6
	fmt.Println(variable6, variable7)
}