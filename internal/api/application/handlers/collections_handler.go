package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	entity_errors "text.io/internal/entities/errors"
	"text.io/internal/entities/models"
	"text.io/internal/infrastructure/services"
)

type CollectionsHandler struct {
	service *services.CollectionsService
}

func NewCollectionsHandler(service *services.CollectionsService) *CollectionsHandler {
	return &CollectionsHandler{
		service: service,
	}
}

func (h *CollectionsHandler) GetCollection(w http.ResponseWriter, r *http.Request) {
	fingerprint := chi.URLParam(r, "fingerprint")

	collection, err := h.service.GetCollection(fingerprint)
	if err != nil {
		http.Error(w, "Collection not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(collection)
}

func (h *CollectionsHandler) ListCollections(w http.ResponseWriter, r *http.Request) {
	collections, err := h.service.ListCollections()
	if err != nil {
		http.Error(w, "Failed to retrieve collections", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(collections)
}

func (h *CollectionsHandler) CreateCollection(w http.ResponseWriter, r *http.Request) {
	var collection models.CreateCollection
	if err := json.NewDecoder(r.Body).Decode(&collection); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.CreateCollection(collection)
	if err != nil {
		switch err {
		case entity_errors.ErrInvalidCollection:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case entity_errors.ErrCollectionExists:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, "Error saving the collection", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(collection)
}
