package database

import (
	"database/sql"
	"fmt"

	"github.com/antonovegorv/orders-service/config"
	_ "github.com/lib/pq"
)

// A variable to store database connection.
var DB *sql.DB

// Connects to the Database using environmental variables.
func Connect() error {
	var err error

	// Make a connection string.
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Get("DB_HOST"),
		config.Get("DB_PORT"),
		config.Get("DB_USER"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_NAME"),
	)

	// Open a connectrion.
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	return nil
}
