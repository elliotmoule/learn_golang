package main

import "fmt"

// ########### OPERATORS ###########

// Operators demonstrates which operators are available within Go, and how they can be used.
func Operators() {
	PrintHeading("operators:")
	x, y := 5, 6

	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x mod y = ", x%y)
}
