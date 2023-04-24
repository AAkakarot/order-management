package mysql

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	DB *sql.DB
)

func GetDB() (*sql.DB, error) {
	if DB == nil {
		// Initialize a new connection to the database
		db, err := sql.Open("mysql", fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		))
		if err != nil {
			return nil, fmt.Errorf("error opening database connection: %v", err)
		}
		DB = db
	}

	return DB, nil
}
