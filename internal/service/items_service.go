package service

import (
	"errors"

	"text.io/internal/database"
	"text.io/internal/models"
)

var (
	ErrInvalidItem = errors.New("invalid item: ID and Name are required")
	ErrItemExists  = errors.New("item already exists")
)

type ItemsService struct {
	repo database.ItemRepository
}

func NewService(repo database.ItemRepository) *ItemsService {
	return &ItemsService{
		repo: repo,
	}
}

func (s *ItemsService) CreateItem(item models.Item) error {
	// Validate the item
	if item.ID == "" || item.Name == "" {
		return ErrInvalidItem
	}

	existingItem, err := s.repo.GetItem(item.ID)
	if err != nil && existingItem.ID != "" {
		return ErrItemExists
	}

	return s.repo.CreateItem(item)
}

func (s *ItemsService) GetItem(id string) (models.Item, error) {
	if id == "" {
		return models.Item{}, errors.New("item ID is required")
	}

	return s.repo.GetItem(id)
}

func (s *ItemsService) ListItems() ([]models.Item, error) {
	return s.repo.ListItems()
}
