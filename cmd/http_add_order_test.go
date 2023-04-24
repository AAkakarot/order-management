package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"service/mysql"
)

func TestCreateOrderHandler(t *testing.T) {
	// Create a new HTTP request
	reqBody := mysql.Order{
		ID:           "abcdef-123456",
		Status:       "PENDING_INVOICE",
		Items:        []mysql.OrderItem{{ID: "123456", Description: "a product description", Price: 12.40, Quantity: 1}},
		Total:        12.40,
		CurrencyUnit: "USD",
	}
	reqBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "api/v1/orders", bytes.NewBuffer(reqBytes))

	// Create a new recorder that will record the response from the HTTP handler
	rr := httptest.NewRecorder()

	// Call the HTTP handler
	handler := http.HandlerFunc(addOrderHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if rr.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusCreated)
	}

	// Check the response body
	expected := `{"message":"Order created successfully","order":{"id":"abcdef-123456","status":"PENDING_INVOICE","items":[{"id":"123456","description":"a product description","price":12.4,"quantity":1}],"total":12.4,"currencyUnit":"USD"}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
