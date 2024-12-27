package main

import (
	"REST-server/models"
	"REST-server/routes"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleRequests() {
	port := 8080
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/api/orders", routes.CreateOrder).Methods("POST")
	myRouter.HandleFunc("/api/orders/{id}", routes.ReadOrder).Methods("GET")
	myRouter.HandleFunc("/api/orders/{id}", routes.UpdateOrder).Methods("PUT")
	myRouter.HandleFunc("/api/orders/{id}", routes.DeleteOrder).Methods("DELETE")

	fmt.Printf("Server is running on http://localhost:%d", port)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", port), myRouter)
}

func main() {
	models.ConnectToDB()
	handleRequests()
}
