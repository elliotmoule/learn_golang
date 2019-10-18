package main

import "fmt"

// PrintHeading provides an easy way to print a heading which is encapsulated with a block of hashes
func PrintHeading(name string) {
	var hashes string = "################"
	fmt.Println("\n" + hashes + "\n## " + name + "\n" + hashes)
}
