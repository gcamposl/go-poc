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
	Area() float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
