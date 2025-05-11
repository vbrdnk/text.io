package database

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"text.io/configs"
	"text.io/ent"
)

// Init DB connection
func InitDB(config configs.Config) (error, *ent.Client) {
	// Create connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	client, err := ent.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return nil, client
}
