package main

import "fmt"

// ########### LOOPS ###########

// Looping covers how loops work in go, and provides multiple patterns as examples.
func Looping() {
	PrintHeading("looping")
	limit := 5
	fmt.Println("Pattern 1:")
	for i := 0; i < limit; i++ {
		for j := 0; j < i; j++ {
			fmt.Println("*")
		}
		fmt.Println()
	}

	fmt.Println("Pattern 2:")
	for i := 0; i < limit; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
