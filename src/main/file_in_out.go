package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ########### FILE I/O ###########

// FileInOut covers how to handle file i/o in Go (with streams included)
func FileInOut() {
	PrintHeading("file i/o")
	fileLoc := "sample.txt"
	file, err := os.Create(fileLoc)

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Hello, nice to meet you. This file was created using GoLang")
	file.Close()

	stream, err := ioutil.ReadFile(fileLoc)

	if err != nil {
		log.Fatal(err)
	}
	s1 := string(stream)
	fmt.Println("Outputting the contents of a file:", s1)
}
