package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"text.io/internal/database"
	"text.io/internal/models"
)

func getItemByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var item models.Item
	err := database.DB.QueryRow("SELECT id, name, created_at FROM items WHERE id = $1", id).Scan(&item.ID, &item.Name, &item.CreatedAt)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func listItems(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name, created_at FROM items")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	// Parse rows into items
	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.CreatedAt); err != nil {
			http.Error(w, "Error reading data", http.StatusInternalServerError)
			return
		}

		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
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

	_, err := database.DB.Exec("INSERT INTO items (id, name) VALUES ($1, $2)", item.ID, item.Name)
	if err != nil {
		http.Error(w, "Error saving the item", http.StatusInternalServerError)
		return
	}

	item.CreatedAt = time.Now()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
