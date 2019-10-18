package main

import "fmt"

// ########### DEFER RECOVER PANIC ###########

// DeferRecoverPanic covers the the basics for these Go specific statements
func DeferRecoverPanic() {
	PrintHeading("control structure")

	fmt.Println("Defer in use:") // Defer running a line until the end of a function.
	deferInUse()

	fmt.Println("Recover in use:") // Recover from a Panic (Exception)
	recoverInUse()

	fmt.Println("Panic in use:") // Panic is essentially an exception statement, providing when things have gone wrong.
	panicInUse()
}

func deferInUse() {
	defer firstRun()
	secondRun()
}

func firstRun() {
	fmt.Println("I executed first")
}

func secondRun() {
	fmt.Println("I executed second")
}

func recoverInUse() {
	fmt.Println(div(3, 0))
	fmt.Println(div(3, 5))
}

func div(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	solution := num1 / num2
	return solution
}

func panicInUse() {
	defer func() {
		fmt.Println(recover())
	}()
	panic("!!Panic!!")
}
