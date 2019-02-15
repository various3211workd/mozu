package server

import (
	"fmt"
	"net/http"
)

/*
	handler function
*/
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}

/*
	Run funciton
*/
func Run() {
	port := 8080
	fmt.Printf("Starting server at Port %d", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
