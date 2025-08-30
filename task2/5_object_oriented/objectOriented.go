package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width float64
	hight float64
}

func (r Rectangle) area() float64 {
	return r.width * r.hight
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.hight)
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	rect := Rectangle{
		width: 5,
		hight: 3,
	}

	cir := Circle{
		radius: 4,
	}

	shapes := []Shape{
		rect,
		cir,
	}

	for _, shape := range shapes {
		fmt.Printf("shape area: %.2f \t", shape.area())
		fmt.Printf("shape perimeter: %.2f\n", shape.perimeter())
	}

}
