package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateOrderStatusHandler(t *testing.T) {
	// Create a new HTTP request
	reqBody := UpdateOrderStatusRequest{
		Status: "SHIPPED",
	}
	reqBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("PUT", "api/v1/orders/abcdef-123456/status", bytes.NewBuffer(reqBytes))

	// Create a new recorder that will record the response from the HTTP handler
	rr := httptest.NewRecorder()

	// Call the HTTP handler
	handler := http.HandlerFunc(updateOrderStatusHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	// Check the response body
	expected := `{"message":"Order status updated successfully","order":{"id":"abcdef-123456","status":"SHIPPED","items":[{"id":"123456","description":"a product description","price":12.4,"quantity":1}],"total":12.4,"currencyUnit":"USD"}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
