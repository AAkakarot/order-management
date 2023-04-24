# order-management

# Order Management Service

This project is an order management service built using Go programming language and MySQL database. The service supports adding an order with different status values and updating the status on the existing order. Additionally, the service provides APIs to fetch orders based on all the fields of the order in a sorted and filtered way.

Order Payload
The order payload is in the JSON format and consists of the following fields:

id (string): a unique identifier for the order
status (string): the current status of the order, which can be one of the following: PENDING_INVOICE, INVOICED, SHIPPED, or CANCELLED
items (array of objects): a list of items in the order
total (float): the total price of the order
currencyUnit (string): the currency unit of the order, such as "USD"
Each item in the items array consists of the following fields:

id (string): a unique identifier for the item
description (string): a brief description of the item
price (float): the price of the item
quantity (integer): the quantity of the item
API Endpoints
The service provides the following REST API endpoints:

POST /orders: Add a new order to the system
PUT /orders/:id/status: Update the status of an existing order
GET /orders: Fetch all orders with the ability to filter and sort the results based on the order payload fields
Stack
The order management service is built using the following stack:

Go programming language
MySQL database
JSON - HTTP REST APIs
How to Run the Service
To run the service, follow these steps:

Clone the repository from GitHub
Set up a MySQL database and update the configuration details in the config/config.go file
Install Go and other dependencies
Run the service using the go run main.go command
Testing
The service is accompanied by tests, which can be run using the go test command. Tests are located in the test directory and cover different scenarios for adding orders, updating orders, and fetching orders.

Documentation
Documentation for the service can be found in the docs directory. This includes a detailed description of the API endpoints, the order payload, and the database schema. Additionally, the README.md file provides information on how to run the service and run tests.

Containerization
The service has been containerized using Docker. The Dockerfile is located in the root directory of the repository. To build the Docker image, run the following command:
docker build -t order-management-service .
