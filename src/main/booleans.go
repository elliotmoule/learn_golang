package main

import "fmt"

// ########### BOOLEANS ###########

// Booleans covers how bools are used within Go. This provides available values, and logical operators.
func Booleans() {
	PrintHeading("booleans")
	isBool := true
	hate := false

	fmt.Println(isBool && hate)
	fmt.Println(isBool || hate)
	fmt.Println(!isBool)
}
