package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

// ########### WEB SERVER ###########

// WebServer demosntrates how Go can create simplistic web servers incredibly easily.
func WebServer() {
	PrintHeading("web server")
	url := "http://localhost:8080/"
	http.HandleFunc("/", handler1)
	http.HandleFunc("/Hello", handler2)
	fmt.Println("Opening Web Url:", url)
	err1 := exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	err2 := exec.Command("rundll32", "url.dll,FileProtocolHandler", url+"Hello").Start()
	http.ListenAndServe(":8080", nil)

	if err1 != nil || err2 != nil {
		log.Fatal()
	}
}

func handler1(w http.ResponseWriter, r *http.Request) {
	message := "Welcome to my Home Page!\n"
	fmt.Fprintf(w, message)
	fmt.Println("Web Server 1 - Running!\nMessage:\n'" + message + "'")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	message := "Hello World!"
	fmt.Fprintf(w, message)
	fmt.Println("Web Server 2 - Running!\nMessage:\n'" + message + "'")
}
