package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"text.io/configs"
	"text.io/internal/database"
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

	handler := NewHandler(repo)

	// Add built-in middleware
	server.router.Use(middleware.Logger)
	server.router.Use(middleware.Recoverer)
	server.router.Use(middleware.Timeout(time.Duration(config.Timeout) * time.Second))

	server.router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Define the routes
	server.router.Route("/api", func(r chi.Router) {
		r.Get("/items", handler.ListItems)
		r.Post("/items", handler.CreateItem)

		// Path parameter example
		r.Get("/items/{id}", handler.GetItem)
	})

	return server
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.config.Port)
	fmt.Printf("Server starting on port %s...\n", addr)
	return http.ListenAndServe(addr, s.router)
}
