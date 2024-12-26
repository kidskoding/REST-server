package main

import (
	"fmt"
	"net/http"
)

func main() {
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page!")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	port := 8080
	fmt.Println("Server is running on port", port)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
}
