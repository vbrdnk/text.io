package server

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

func Run() {
	config := configs.LoadConfig()

	// Initialize the database
	if err := database.InitDB(config); err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	defer database.CloseDB()

	repo := database.NewPostgresItemRepository(database.DB)

	server := api.NewServer(config, repo)
	go func() {
		if err := server.Start(); err != nil {
			log.Fatalf("Server error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")
}
