package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"text.io/internal/database"
	"text.io/internal/models"
)

type Handler struct {
	repo database.ItemRepository
}

func NewHandler(repo database.ItemRepository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) GetItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	item, err := h.repo.GetItem(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (h *Handler) ListItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.repo.ListItems()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the item
	if item.ID == "" || item.Name == "" {
		http.Error(w, "ID and Name are required", http.StatusBadRequest)
		return
	}

	err := h.repo.CreateItem(item)
	if err != nil {
		http.Error(w, "Error saving the item", http.StatusInternalServerError)
		return
	}

	item.CreatedAt = time.Now()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
