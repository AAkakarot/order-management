package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"service/mysql"
)

func addOrderHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the order payload from the request body
	var order mysql.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error parsing order payload: %v", err)
		return
	}

	// Validate the order data
	if order.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Order ID is required")
		return
	}

	// Save the order to the database
	err = order.Save()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error saving order: %v", err)
		return
	}

	// Respond with the saved order as JSON
	w.Header().Set("Content-Type", "application/json")
	log.Println("added order with id: ", order.ID)
	json.NewEncoder(w).Encode(order)
}
