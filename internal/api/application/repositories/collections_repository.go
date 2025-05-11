package repositories

import "text.io/internal/entities/models"

type CollectionsRepository interface {
	GetCollection(fingerprint string) (models.Collection, error)
	CreateCollection(shortUrl string, collection models.CreateCollection) error
	ListCollections() ([]models.Collection, error)
}
