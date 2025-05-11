package services

import (
	"text.io/ent"
	"text.io/internal/api/application/repositories"
	entity_errors "text.io/internal/entities/errors"
	"text.io/internal/entities/models"

	"text.io/pkg"
)

type CollectionsService struct {
	repo repositories.CollectionsRepository
}

func NewCollectionsService(repo repositories.CollectionsRepository) *CollectionsService {
	return &CollectionsService{
		repo: repo,
	}
}

func (s *CollectionsService) CreateCollection(collection models.CreateCollection) error {
	// Validate the collection
	if collection.Title == "" || collection.Description == "" {
		return entity_errors.ErrInvalidCollection
	}

	shortUrl := pkg.GenerateURLFingerprint(collection.Title)

	exists, err := s.repo.CheckIfCollectionExists(shortUrl)
	if err != nil || exists {
		return entity_errors.ErrCollectionExists
	}

	return s.repo.CreateCollection(shortUrl, collection)
}

func (s *CollectionsService) GetCollection(fingerprint string) (*ent.Collection, error) {
	return s.repo.GetCollection(fingerprint)
}

func (s *CollectionsService) ListCollections() ([]*ent.Collection, error) {
	return s.repo.ListCollections()
}
