package main

import (
	"fmt"
	"net/http"
)

//const to store port number
const PortNumber := 8080

func Home(w http.ResponseWriter, r *http.Response) {
	fmt.Fprintf(w, "The Golang home page")
}

func About(w http.ResponseWriter, r *http.Response) {
	fmt.Fprintf(w, "The Golang about page")
}


//Simple main function to test out the output
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/About", About)

	//open the port to listen
	_ = http.ListenAndServe(PortNumber, nil)
}