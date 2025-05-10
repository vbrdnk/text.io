package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"text.io/configs"
	"text.io/internal/api/application/handlers"
	"text.io/internal/api/application/repositories"
	"text.io/internal/infrastructure/services"
)

type Server struct {
	router chi.Router
	config configs.Config
}

func NewServer(config configs.Config, repo repositories.LinkRepository) *Server {
	server := &Server{
		router: chi.NewRouter(),
		config: config,
	}

	linksService := services.NewService(repo)
	linksHandler := handlers.NewHandler(linksService)

	// Add built-in middleware
	server.router.Use(middleware.Logger)
	server.router.Use(middleware.Recoverer)
	server.router.Use(middleware.Timeout(time.Duration(config.Timeout) * time.Second))

	// Define the routes
	server.router.Route("/api", func(r chi.Router) {
		r.Route("/links", func(r chi.Router) {
			r.Get("/", linksHandler.ListLinks)
			r.Post("/", linksHandler.CreateLink)
			r.Get("/{fingerprint}", linksHandler.GetLink)
		})
	})

	return server
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.config.Port)
	fmt.Printf("Server starting on port %s...\n", addr)
	return http.ListenAndServe(addr, s.router)
}
