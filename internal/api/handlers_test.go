package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"text.io/internal/database"
	"text.io/internal/models"
)

func TestGetItem(t *testing.T) {
	// Create a mock repository
	mockRepo := database.NewMockItemRepository()

	// Add a test item to the mock repository
	testItem := models.Item{ID: "test1", Name: "Test Item"}
	mockRepo.CreateItem(testItem)

	// Create a handler with the mock repository
	handler := NewHandler(mockRepo)

	// Create a test request
	req, err := http.NewRequest("GET", "/api/items/test1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Add URL parameters
	// If using chi router:
	chiCtx := chi.NewRouteContext()
	chiCtx.URLParams.Add("id", "test1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

	// Or for simpler testing, just add a query parameter:
	// q := req.URL.Query()
	// q.Add("id", "test1")
	// req.URL.RawQuery = q.Encode()

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler.GetItem(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Parse response
	var responseItem models.Item
	if err := json.Unmarshal(rr.Body.Bytes(), &responseItem); err != nil {
		t.Errorf("couldn't parse response: %v", err)
	}

	// Check response data
	if responseItem.ID != testItem.ID || responseItem.Name != testItem.Name {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseItem, testItem)
	}
}

func TestCreateItem(t *testing.T) {
	// Create a mock repository
	mockRepo := database.NewMockItemRepository()

	// Create a handler with the mock repository
	handler := NewHandler(mockRepo)

	// Create a test request with JSON body
	reqBody := `{"id":"test2","name":"New Test Item"}`
	req, err := http.NewRequest("POST", "/api/items", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler.CreateItem(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Verify item was created in repository
	item, err := mockRepo.GetItem("test2")
	if err != nil {
		t.Errorf("item was not created in repository: %v", err)
	}

	if item.Name != "New Test Item" {
		t.Errorf("item has wrong name: got %v want %v",
			item.Name, "New Test Item")
	}
}

func TestListItems(t *testing.T) {
	// Create a mock repository
	mockRepo := database.NewMockItemRepository()

	// Add test items
	mockRepo.CreateItem(models.Item{ID: "1", Name: "Item 1"})
	mockRepo.CreateItem(models.Item{ID: "2", Name: "Item 2"})

	// Create a handler with the mock repository
	handler := NewHandler(mockRepo)

	// Create a test request
	req, err := http.NewRequest("GET", "/api/items", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler.ListItems(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Parse response
	var items []models.Item
	if err := json.Unmarshal(rr.Body.Bytes(), &items); err != nil {
		t.Errorf("couldn't parse response: %v", err)
	}

	// Check response data
	if len(items) != 2 {
		t.Errorf("expected 2 items, got %d", len(items))
	}

	// You could also check for specific items if needed
	hasItem1 := false
	hasItem2 := false
	for _, item := range items {
		if item.ID == "1" && item.Name == "Item 1" {
			hasItem1 = true
		}
		if item.ID == "2" && item.Name == "Item 2" {
			hasItem2 = true
		}
	}

	if !hasItem1 || !hasItem2 {
		t.Errorf("response missing expected items")
	}
}
