package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"text.io/configs"
	"text.io/ent"
	"text.io/internal/api/application/handlers"
	"text.io/internal/api/application/repositories"
	"text.io/internal/infrastructure/services"
)

type Server struct {
	router chi.Router
	config configs.Config
}

func NewServer(
	config configs.Config,
	client *ent.Client,
	linksRepo repositories.LinkRepository,
	collectionsRepo repositories.CollectionsRepository,
) *Server {
	server := &Server{
		router: chi.NewRouter(),
		config: config,
	}

	// Initialize all services
	linksService := services.NewLinksService(linksRepo)
	collectionsService := services.NewCollectionsService(collectionsRepo)

	// Initialize all handlers
	linksHandler := handlers.NewLinksHandler(linksService)
	collectionsHandler := handlers.NewCollectionsHandler(collectionsService)

	// Add built-in middleware
	server.router.Use(middleware.Logger)
	server.router.Use(middleware.Recoverer)
	server.router.Use(middleware.Timeout(time.Duration(config.Timeout) * time.Second))

	// Define the routes
	server.router.Route("/api", func(r chi.Router) {
		// Links routes
		r.Route("/links", func(r chi.Router) {
			r.Get("/", linksHandler.ListLinks)
			r.Post("/", linksHandler.CreateLink)
			r.Get("/{fingerprint}", linksHandler.GetLink)
		})

		// Collections routes
		r.Route("/collections", func(r chi.Router) {
			r.Get("/", collectionsHandler.ListCollections)
			r.Post("/", collectionsHandler.CreateCollection)
			r.Get("/{fingerprint}", collectionsHandler.GetCollection)
		})
	})

	return server
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.config.Port)
	fmt.Printf("Server starting on port %s...\n", addr)
	return http.ListenAndServe(addr, s.router)
}
