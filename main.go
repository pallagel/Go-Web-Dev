package main

import (
	"fmt"
	"net/http"
)

//Simple main function to test out the output
func main() {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		_, err := fmt.Fprintf(w, "Hello to Go Web World!!")
		if err != nil {
			fmt.Println("An error occured: ", err)
		}
	})

	//open the port to listen
	_ = http.ListenAndServe(":8080", nil)
}