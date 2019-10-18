package main

import "fmt"

// ########### BASICS ###########

// Basics covers primitive/ scalar variables which are available to Go.
func Basics() {
	PrintHeading("basics")
	var a int = 5 // Longhand
	var b float32 = 4.32
	const pi float64 = 3.1415139475

	x, y := 14, 15 // Shorthand

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(pi)
	fmt.Println(x, ",", y)
}
