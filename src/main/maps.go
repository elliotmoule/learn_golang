package main

import "fmt"

// ########### MAPS ###########

// Maps work like a Dictionary, utilising Key Value Pairs.
func Maps() {
	PrintHeading("maps")
	StudentAge := make(map[string]int)
	StudentAge["Aryya"] = 23
	StudentAge["Elliot"] = 30
	StudentAge["Jim"] = 21
	StudentAge["Cathy"] = 34
	StudentAge["Olivia"] = 27
	StudentAge["Claire"] = 22

	fmt.Println("Print All:\n", StudentAge)

	// Keys are case-specific
	fmt.Println("Print One: (Claire) ==", StudentAge["Claire"])

	fmt.Println("Length of Map:", len(StudentAge))

	fmt.Println("Maps within Maps:")
	// Commas needed between all fields
	superhero :=
		map[string]map[string]string{
			"Superman": map[string]string{
				"RealName": "Clark Kent",
				"City":     "Metropolis",
			},
			"Batman": map[string]string{
				"RealName": "Bruce Wayne",
				"City":     "Gotham City",
			},
		}

	if temp, hero := superhero["Superman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])
	}
}
