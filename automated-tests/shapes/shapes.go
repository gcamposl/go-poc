package shapes

import (
	"math"
)

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	area() float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
