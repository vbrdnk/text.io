package repositories

import (
	"text.io/ent"
	"text.io/internal/entities/models"
)

type CollectionsRepository interface {
	CheckIfCollectionExists(fingerprint string) (bool, error)
	GetCollection(fingerprint string) (*ent.Collection, error)
	CreateCollection(shortUrl string, collection models.CreateCollection) error
	ListCollections() ([]*ent.Collection, error)
}
