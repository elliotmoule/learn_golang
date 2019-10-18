package main

import "fmt"

// ########### FUNCTIONS ###########

// Functions covers how a simple function works within Go.
func Functions() {
	PrintHeading("functions")

	x, y := 5, 6

	fmt.Println("Adding function:", add(x, y))
}

func add(num1, num2 int) int {
	return num1 + num2
}
