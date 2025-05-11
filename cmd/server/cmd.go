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
	"text.io/internal/infrastructure/repositories"
)

func Run() {
	config := configs.LoadConfig()

	// Initialize the database
	err, client := database.InitDB(config)
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	defer client.Close()

	linksRepo := repositories.NewPostgresLinkRepository(client)
	collectionsRepo := repositories.NewPostgresCollectionRepository(client)

	server := api.NewServer(config, client, linksRepo, collectionsRepo)
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
