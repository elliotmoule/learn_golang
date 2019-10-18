package main

import (
	"fmt"
	"math"
)

// ########### INTERFACES ###########

// Interfaces covers how to create a simple interface, wherein classes have a single relation
func Interfaces() {
	PrintHeading("interfaces")
	rect := Rectangle{50, 60}
	circ := Circle{7}

	fmt.Println("Area of Rectangle is:", getArea(rect)) // This rect is based upon 'Rectangle' from the 'structures.go' file
	fmt.Println("Area of circle is:", getArea(circ))
}

// Shape is the root relation for Rectangle and Circle
type Shape interface {
	area() float64
}

// Circle is a Shape based structure.
type Circle struct {
	radius float64
}

func (r1 Rectangle) area() float64 {
	return r1.height * r1.width
}
func (c1 Circle) area() float64 {
	return math.Pi * math.Pow(c1.radius, 2)
}
func getArea(shape Shape) float64 {
	return shape.area()
}
