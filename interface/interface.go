package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	height float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func writeArea(s shape) {
	fmt.Printf("the area of the shape is %0.2f\n", s.area())
}

func main() {
	fmt.Println("interfaces")
	c := circle{2}
	r := rectangle{2, 4}
	writeArea(c)
	writeArea(r)
}
