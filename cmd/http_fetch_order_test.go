func TestGetOrdersHandler(t *testing.T) {
	// Create a request with query parameters
	req, err := http.NewRequest("GET", "api/v1/orders?status=PENDING_INVOICE&from=2022-01-01&to=2022-12-31&sortBy=id&sortOrder=asc", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the GetOrdersHandler function with the request and response recorder
	fetchOrdersHandler(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	// Parse the response body
	var responseBody struct {
		Message string  `json:"message"`
		Orders  []Order `json:"orders"`
	}
	if err := json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
		t.Fatal(err)
	}

	// Check the response message
	expectedMessage := "Successfully retrieved orders"
	if responseBody.Message != expectedMessage {
		t.Errorf("Expected message %q, got %q", expectedMessage, responseBody.Message)
	}

	// Check the number of orders in the response
	expectedNumOrders := 1
	if len(responseBody.Orders) != expectedNumOrders {
		t.Errorf("Expected %d orders, got %d", expectedNumOrders, len(responseBody.Orders))
	}

