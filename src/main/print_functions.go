package main

import "fmt"

// ########### PRINT FUNCTIONS ###########

// PrintFunctions demonstrates the various print functions available within Go.
func PrintFunctions() {
	PrintHeading("print functions")
	var name string = "Jamie Johnson"
	const pi float64 = 3.1412435
	win := true
	x := 5

	fmt.Println("Get string length: ", len(name))
	fmt.Println("Concat strings: ", name+" is a good guy!")

	fmt.Printf("Print a float: %f \n", pi)
	fmt.Printf("Print float with precision: %.3f \n", pi)
	fmt.Printf("Print a Var Type: %T \n", name)
	fmt.Printf("Print boolean: %t \n", win)
	fmt.Printf("Print integer: %d \n", x)
	fmt.Printf("Print the binary of: %b \n", 25)
	fmt.Printf("Print an ASCII character: 33 == '%c' \n", 33)
	fmt.Printf("Print a hex code: 15 == '%x' \n", 15)
	fmt.Printf("Print scientific notation: pi == '%e' \n", pi)
}
