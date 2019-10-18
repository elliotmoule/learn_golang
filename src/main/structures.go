package main

import "fmt"

// ########### STRUCTURES ###########

// Structures demonstrates how structures can be used and work within Go.
func Structures() {
	PrintHeading("structures")
	rect1 := Rectangle{height: 10, width: 5}
	fmt.Println("Explicitly set struct vals:", rect1)

	rect2 := Rectangle{10, 5}
	fmt.Println("Lazy set struct vals:", rect2)

	fmt.Println("Extending struct with custom function:", rect1.pArea())
}

// Rectangle is a structure with a height and width.
type Rectangle struct {
	height float64
	width  float64
}

// Area, using pointer.
func (rect *Rectangle) pArea() float64 {
	return rect.height * rect.width
}
