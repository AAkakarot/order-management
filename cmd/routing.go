package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	// Create a new router.
	r := mux.NewRouter()

	// Register the handlers.

	// Homepage
	r.HandleFunc("/api/v1", HomePage).Methods("GET")

	// Get Order
	r.HandleFunc("/api/v1/orders/{id}", fetchOrdersHandler).Methods("GET")

	// Add Order
	r.HandleFunc("/api/v1/orders", addOrderHandler).Methods("POST")

	// Update Order
	r.HandleFunc("/api/v1/orders/{id}", updateOrderStatusHandler).Methods("PUT")

	// TODO: List of Orders
	//r.HandleFunc("/api/v1/orders", ListOrdersHandler).Methods("GET")

	// Start the server.
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
