package main

import (
	"fmt"
	"net/http"
	"github.com/gowebdev/pkg/handlers"
	"log"
)

//
const PortNumber = ":4444"

// main is the main function execute in the program
func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)

	fmt.Println(fmt.Sprintf("Staring application on port %s", PortNumber))

	//if port cant be open
	err := http.ListenAndServe(PortNumber, nil)

	//log the error and exit the program
	if err != nil {
		log.Fatal("Unable to start the server on port: ", PortNumber, "And error is :", err)
	}
}
