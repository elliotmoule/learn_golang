package main

import "fmt"

// ########### ARRAYS ###########

// Arrays covers how Go can utilise arrays. This includes initialisation, setting, splitting, copying, and appending.
func Arrays() {
	PrintHeading("arrays")
	var EvenNum [5]int
	EvenNumAlt := [5]int{0, 2, 4, 6, 8} // Alternate initialisation of an array.

	// Initialise Array via Index
	EvenNum[0] = 0
	EvenNum[1] = 2
	EvenNum[2] = 4
	EvenNum[3] = 6
	EvenNum[4] = 8

	fmt.Println("Print index initialised array value.")
	fmt.Println(EvenNum[2])
	fmt.Println("Print group initialised array value.")
	fmt.Println(EvenNumAlt[2])

	// Can use underscore if index is not needed.
	fmt.Println("Array print from loop.")
	for _, value := range EvenNum {
		fmt.Println(value)
	}

	fmt.Println("Sliced Arrays:")
	numSlice := []int{5, 4, 3, 2, 1}

	fmt.Println("Initial slice: ", numSlice)

	sliced := numSlice[3:5]
	fmt.Println("Range Slice (3-5): ", sliced)

	slice2 := make([]int, 5, 10) // Create an empty slice and define the data type.

	copy(slice2, numSlice) // Copy contents of one slice to another
	fmt.Println("Copy slice: ", slice2)

	slice3 := append(numSlice, 3, 9, -1)
	fmt.Println("Append Slice: ", slice3)
}
