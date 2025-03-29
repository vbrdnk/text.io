package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"text.io/configs"
)

// DB is the database connection
var DB *sql.DB

// Init DB connection
func InitDB(config configs.Config) error {
	// Create connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	// Open connection
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	// Check connection
	if err = DB.Ping(); err != nil {
		return err
	}

	// Init database tables
	if err = createTable(); err != nil {
		return err
	}

	log.Println("Database connection established")
	return nil
}

func createTable() error {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS items (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW()
		)
	`)

	return err
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
