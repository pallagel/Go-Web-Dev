package main

import (
	"fmt"
	"net/http"
)

//const to store port number
const PortNumber = ":8080"

//Home - Direct to home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The Golang home page")
}

//About - About page function
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The Golang about page")
}


//Simple main function to test out the output
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	//open the port to listen
	_ = http.ListenAndServe(PortNumber, nil)
}