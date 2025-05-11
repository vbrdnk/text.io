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
	// DB.Exec(`
	// 	DROP TABLE links
	// `)
	//
	// DB.Exec(`
	// 	DROP TABLE collections
	// `)

	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS links (
			fingerprint TEXT PRIMARY KEY UNIQUE,
			label TEXT,
			url TEXT NOT NULL,
			created_by TEXT,
			created_at TIMESTAMP NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS collections (
			fingerprint TEXT PRIMARY KEY UNIQUE,
			title TEXT NOT NULL,
			description TEXT,
			published BOOLEAN DEFAULT FALSE,
			created_by TEXT,
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
