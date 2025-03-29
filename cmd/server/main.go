package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"text.io/configs"
	"text.io/internal/api"
	"text.io/internal/database"
)

func main() {
	config := configs.LoadConfig()

	// Initialize the database
	if err := database.InitDB(config); err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	defer database.CloseDB()

	go func() {
		if err := api.RunServer(config); err != nil {
			log.Fatalf("Server error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")
}
