package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"text.io/configs"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func RunServer(config configs.Config) error {
	// Create a new chi router
	r := chi.NewRouter()

	// Add built-in middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Duration(config.Timeout) * time.Second))

	r.Get("/hello", helloHandler)

	// Define the routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/items", listItems)
		r.Post("/items", createItem)

		// Path parameter example
		r.Get("/items/{id}", getItemByID)
	})

	// Start the server on port 8080
	fmt.Printf("Server starting on port %d...\n", config.Port)

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)
}
