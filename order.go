package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type Order struct {
	ID           string       `json:"id"`
	Status       string       `json:"status"`
	Items        []OrderItem  `json:"items"`
	Total        float64      `json:"total"`
	CurrencyUnit string       `json:"currency_unit"`
	CreatedAt    sql.NullTime `json:"created_at"`
}

type OrderItem struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type OrderQuery struct {
	ID           string
	Status       string
	Total        string
	CurrencyUnit string
	SortBy       string
	SortOrder    string
}

func (o *Order) Save() error {
	db, err := GetDB()
	if err != nil {
		return fmt.Errorf("error getting database connection: %v", err)
	}
	defer db.Close()

	now := time.Now().UTC()
	createdAt := sql.NullTime{Time: now, Valid: true}

	// Insert the order into the database
	_, err = db.Exec(
		"INSERT INTO orders (id, status, items, total, currency_unit, created_at) VALUES (?, ?, ?, ?, ?, ?)",
		o.ID, o.Status, toJSON(o.Items), o.Total, o.CurrencyUnit, createdAt,
	)
	if err != nil {
		return fmt.Errorf("error inserting order into database: %v", err)
	}

	return nil
}

func (o *Order) UpdateStatus(status string) error {
	db, err := GetDB()
	if err != nil {
		return fmt.Errorf("error getting database connection: %v", err)
	}
	defer db.Close()

	// Update the order status in the database
	_, err = db.Exec("UPDATE orders SET status = ? WHERE id = ?", status, o.ID)
	if err != nil {
		return fmt.Errorf("error updating order status in database: %v", err)
	}

	// Update the order struct
	o.Status = status

	return nil
}

func FetchOrders(query OrderQuery) ([]Order, error) {
	db, err := GetDB()
	if err != nil {
		return nil, fmt.Errorf("error getting database connection: %v", err)
	}
	defer db.Close()

	sqlStmt, args := getSqlQuery(query)
	// Execute the SQL query
	rows, err := db.Query(sqlStmt, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying orders from database: %v", err)
	}
	defer rows.Close()

	// Parse the rows into Order structs
	orders := []Order{}
	for rows.Next() {
		var id string
		var status string
		var itemsJSON string
		var total float64
		var currencyUnit string
		var createdAt sql.NullTime

		err = rows.Scan(&id, &status, &itemsJSON, &total, &currencyUnit, &createdAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning order row: %v", err)
		}

		items := []OrderItem{}
		err = fromJSON(itemsJSON, &items)
		if err != nil {
			return nil, fmt.Errorf("error decoding order items JSON: %v", err)
		}

		order := Order{
			ID:           id,
			Status:       status,
			Items:        items,
			Total:        total,
			CurrencyUnit: currencyUnit,
			CreatedAt:    createdAt,
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func getSqlQuery(query OrderQuery) (string, []interface{}) {
	// Construct the SQL query based on the query parameters
	sqlStmt := "SELECT id, status, items, total, currency_unit, created_at FROM orders"
	args := []interface{}{}
	if query.ID != "" {
		sqlStmt += " WHERE id = ?"
		args = append(args, query.ID)
	}
	if query.Status != "" {
		if len(args) == 0 {
			sqlStmt += " WHERE status = ?"
		} else {
			sqlStmt += " AND status = ?"
		}
		args = append(args, query.Status)
	}
	if query.CurrencyUnit != "" {
		if len(args) == 0 {
			sqlStmt += " WHERE currency_unit = ?"
		} else {
			sqlStmt += " AND currency_unit = ?"
		}
		args = append(args, query.CurrencyUnit)
	}
	if query.Total != "" {
		if len(args) == 0 {
			sqlStmt += " WHERE total <= ?"
		} else {
			sqlStmt += " AND total <= ?"
		}
		args = append(args, query.Total)
	}

	// add sorting
	if query.SortBy != "" {
		sqlStmt += fmt.Sprintf(" ORDER BY %s", query.SortBy)
		if query.SortOrder != "" {
			sqlStmt += fmt.Sprintf(" %s", query.SortOrder)
		}
	}
	return sqlStmt, args
}

func toJSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func fromJSON(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}
