package main

import "fmt"

// ########### RECURSION ###########

// Recursion demonstrates how a recursive pattern can be used within Go.
func Recursion() {
	PrintHeading("recursion")
	num := 5
	fmt.Println("Recursive function:", factorial(num))
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
