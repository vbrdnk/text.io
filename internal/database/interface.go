package database

import "text.io/internal/models"

type ItemRepository interface {
	GetItem(id string) (models.Item, error)
	CreateItem(item models.Item) error
	ListItems() ([]models.Item, error)
}
