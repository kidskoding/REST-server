package main

import (
	"REST-server/models"
	"REST-server/routes"
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
	myRouter.HandleFunc("/api/orders", routes.CreateOrder).Methods("POST")

	fmt.Printf("Server is running on http://localhost:%d", port)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", port), myRouter)
}

func main() {
	models.ConnectToDB()
	handleRequests()
}
