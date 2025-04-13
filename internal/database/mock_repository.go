// internal/database/mock_repository.go
package database

import (
	"database/sql"
	"errors"

	"text.io/internal/models"
)

// MockItemRepository implements ItemRepository for testing
type MockItemRepository struct {
	items map[string]models.Item
}

// NewMockItemRepository creates a new mock repository
func NewMockItemRepository() *MockItemRepository {
	return &MockItemRepository{
		items: make(map[string]models.Item),
	}
}

func (r *MockItemRepository) GetItem(id string) (models.Item, error) {
	item, exists := r.items[id]
	if !exists {
		return models.Item{}, sql.ErrNoRows
	}
	return item, nil
}

func (r *MockItemRepository) CreateItem(item models.Item) error {
	if _, exists := r.items[item.ID]; exists {
		return errors.New("item already exists")
	}
	r.items[item.ID] = item
	return nil
}

func (r *MockItemRepository) ListItems() ([]models.Item, error) {
	items := make([]models.Item, 0, len(r.items))
	for _, item := range r.items {
		items = append(items, item)
	}
	return items, nil
}
