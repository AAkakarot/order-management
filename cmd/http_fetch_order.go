package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"service/mysql"
)

func fetchOrdersHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters
	query := mysql.OrderQuery{
		ID:           r.URL.Query().Get("id"),
		Status:       r.URL.Query().Get("status"),
		Total:        r.URL.Query().Get("total"),
		CurrencyUnit: r.URL.Query().Get("currency_unit"),
		SortBy:       r.URL.Query().Get("sort_by"),
		SortOrder:    r.URL.Query().Get("sort_order"),
	}

	// Fetch the orders from the database
	orders, err := mysql.FetchOrders(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error fetching orders: %v", err)
		return
	}

	// Respond with the fetched orders as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
