package repositories

import (
	"context"
	"fmt"

	"text.io/ent"
	"text.io/ent/collection"
	"text.io/internal/entities/models"
)

type CollectionRepository struct {
	client *ent.Client
}

func NewPostgresCollectionRepository(client *ent.Client) *CollectionRepository {
	return &CollectionRepository{
		client: client,
	}
}

func (r *CollectionRepository) CheckIfCollectionExists(fingerprint string) (bool, error) {
	exists, err := r.client.Collection.Query().Where(collection.Fingerprint(fingerprint)).Exist(context.Background())
	if err != nil {
		return false, fmt.Errorf("failed checking if collection exists: %w", err)
	}

	return exists, nil
}

func (r *CollectionRepository) GetCollection(fingerprint string) (*ent.Collection, error) {
	collection, err := r.client.Collection.
		Query().
		Where(collection.Fingerprint(fingerprint)).
		Only(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed getting link: %w", err)
	}
	return collection, nil
}

func (r *CollectionRepository) CreateCollection(shortUrl string, collection models.CreateCollection) error {
	_, err := r.client.Collection.
		Create().
		SetFingerprint(shortUrl).
		SetTitle("test").
		Save(context.Background())
	if err != nil {
		return fmt.Errorf("failed creating collection %w", err)
	}
	return nil
}

func (r *CollectionRepository) ListCollections() ([]*ent.Collection, error) {
	collections, err := r.client.Collection.
		Query().
		All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed quering the collections: %w", err)
	}
	return collections, nil
}
