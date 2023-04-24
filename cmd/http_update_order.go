package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"service/mysql"
)

func updateOrderStatusHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the order ID and status from the request body
	var data struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error parsing request payload: %v", err)
		return
	}

	// Update the order status in the database
	order := &mysql.Order{ID: data.ID}
	err = order.UpdateStatus(data.Status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error updating order status: %v", err)
		return
	}

	// Respond with the updated order as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
