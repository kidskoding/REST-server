package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "A REST server that tracks customers' pizza orders")
}

func handleRequests() {
	port := 8080
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)

	fmt.Println("Server is running on port", port)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", port), myRouter)
}

func main() {
	connectDB()
	handleRequests()
}
