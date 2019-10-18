package main

import "fmt"

// ########### CONTROL STRUCTURES ###########

// ControlStructures covers If and Switch statement use within Go.
func ControlStructures() {
	PrintHeading("decision making")
	var age int = 18

	fmt.Println("--If Statements:")
	ifStatements(age)
	ifStatements(age - 1)

	fmt.Println("--Switch Statements:")
	switchStatements(age - 4) // 14
	switchStatements(age - 2) // 16
	switchStatements(age)     // 18
	switchStatements(age + 2) // 20
	switchStatements(age + 4) // 22
}

func ifStatements(age int) {
	if age >= 18 {
		fmt.Printf("(%d) Yes, you can vote!\n", age)
	} else {
		fmt.Printf("(%d) No, you cannot vote!\n", age)
	}
}

func switchStatements(age int) {
	var message string = ""

	switch age {
	case 16:
		message = "Start preparing for University."
		break
	case 18:
		message = "You can now vote!"
		break
	case 20:
		message = "Good luck with your University studies."
		break
	default:
		if age > 17 {
			message = "You should have a job"
		} else {
			message = "You're a bit young, maybe focus on your studies."
		}
		break
	}

	fmt.Printf("(%d) "+message+"\n", age)
}
