package main

import "fmt"

// ########### POINTERS ###########

// Pointers demonstrates how pointers (variable memory address') can be used within Go.
func Pointers() {
	PrintHeading("pointers")
	x := 10

	fmt.Println("X currently equal to: ", x)
	changeValueLocal(x) // Won't change the value of x, as will keep scope locally.
	fmt.Println("After local change, X is equal to: ", x)
	changeValue(&x)                                                                            // Will change the value of x.
	fmt.Println("After changing the value of X, by passing it's reference, X now equals: ", x) // Value of 'x'
	fmt.Println("The memory address of X is: ", &x)                                            // Address of 'x'
}

func changeValue(x *int) {
	*x = 7
}

func changeValueLocal(x int) {
	x = 7
}
