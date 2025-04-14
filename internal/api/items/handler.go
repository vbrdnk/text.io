package items

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"text.io/internal/models"
	"text.io/internal/service"
)

type ItemsHandler struct {
	service *service.ItemsService
}

func NewHandler(service *service.ItemsService) *ItemsHandler {
	return &ItemsHandler{
		service: service,
	}
}

func (h *ItemsHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	item, err := h.service.GetItem(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (h *ItemsHandler) ListItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.ListItems()
	if err != nil {
		http.Error(w, "Failed to retrieve items", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *ItemsHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.CreateItem(item)
	if err != nil {
		switch err {
		case service.ErrInvalidItem:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case service.ErrItemExists:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, "Error saving the item", http.StatusInternalServerError)
		}
		return
	}

	item.CreatedAt = time.Now()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
