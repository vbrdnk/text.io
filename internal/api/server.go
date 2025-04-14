package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"text.io/configs"
	"text.io/internal/api/items"
	"text.io/internal/database"
	"text.io/internal/service"
)

type Server struct {
	router chi.Router
	config configs.Config
}

func NewServer(config configs.Config, repo database.ItemRepository) *Server {
	server := &Server{
		router: chi.NewRouter(),
		config: config,
	}

	itemsService := service.NewService(repo)
	itemsHandler := items.NewHandler(itemsService)

	// Add built-in middleware
	server.router.Use(middleware.Logger)
	server.router.Use(middleware.Recoverer)
	server.router.Use(middleware.Timeout(time.Duration(config.Timeout) * time.Second))

	// Define the routes
	server.router.Route("/api", func(r chi.Router) {
		r.Route("/items", func(r chi.Router) {
			r.Get("/", itemsHandler.ListItems)
			r.Post("/", itemsHandler.CreateItem)
			r.Get("/{id}", itemsHandler.GetItem)
		})
	})

	return server
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.config.Port)
	fmt.Printf("Server starting on port %s...\n", addr)
	return http.ListenAndServe(addr, s.router)
}
