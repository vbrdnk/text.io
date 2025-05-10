package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	entity_errors "text.io/internal/entities/errors"
	"text.io/internal/entities/models"
	"text.io/internal/infrastructure/services"
)

type LinksHandler struct {
	service *services.LinksService
}

func NewHandler(service *services.LinksService) *LinksHandler {
	return &LinksHandler{
		service: service,
	}
}

func (h *LinksHandler) GetLink(w http.ResponseWriter, r *http.Request) {
	fingerprint := chi.URLParam(r, "fingerprint")

	link, err := h.service.GetLink(fingerprint)
	if err != nil {
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(link)
}

func (h *LinksHandler) ListLinks(w http.ResponseWriter, r *http.Request) {
	links, err := h.service.ListLinks()
	if err != nil {
		http.Error(w, "Failed to retrieve links", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(links)
}

func (h *LinksHandler) CreateLink(w http.ResponseWriter, r *http.Request) {
	var link models.Link
	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.CreateLink(link)
	if err != nil {
		switch err {
		case entity_errors.ErrInvalidLink:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case entity_errors.ErrLinkExists:
			http.Error(w, err.Error(), http.StatusConflict)
		case entity_errors.ErrInvalidUrl:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, "Error saving the link", http.StatusInternalServerError)
		}
		return
	}

	link.CreatedAt = time.Now()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(link)
}
