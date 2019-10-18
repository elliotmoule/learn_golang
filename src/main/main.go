package main

import (
	// Imports auto removed on save, if they're unused (extension)
	"fmt"
)

func main() {
	var dashes string = "-----------------"
	fmt.Println(dashes + "Starting Main" + dashes)

	Basics()
	Operators()
	Booleans()
	Pointers()
	PrintFunctions()
	Looping()
	ControlStructures()
	Arrays()
	Maps()
	Functions()
	Recursion()
	DeferRecoverPanic()
	UnlimitedParameters()
	Structures()
	Interfaces()
	FileInOut()
	WebServer()

	fmt.Println(dashes + "Completed Main" + dashes)
}
