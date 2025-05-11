package services

import (
	"errors"

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

	existingCollection, err := s.repo.GetCollection(shortUrl)
	// TODO: revisit this logic and handle errors properly when db is empty
	if err != nil && existingCollection.Fingerprint != "" {
		return entity_errors.ErrCollectionExists
	}

	return s.repo.CreateCollection(shortUrl, collection)
}

func (s *CollectionsService) GetCollection(fingerprint string) (models.Collection, error) {
	if fingerprint == "" {
		return models.Collection{}, errors.New("collection Fingerprint is required")
	}

	return s.repo.GetCollection(fingerprint)
}

func (s *CollectionsService) ListCollections() ([]models.Collection, error) {
	return s.repo.ListCollections()
}
