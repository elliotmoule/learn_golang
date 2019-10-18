package main

import "fmt"

// ########### UNLIMITED PARAMETERS ###########

// UnlimitedParameters demonstrates how any number of parameters can be passed to function within Go.
func UnlimitedParameters() {
	PrintHeading("unlimited parameters")
	fmt.Println("Passing five args for summing:", sum(10, 20, 30, 40, 50))
}

func sum(args ...int) int {
	sum := 0
	for _, value := range args {
		sum += value
	}
	return sum
}
