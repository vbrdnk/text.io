package repositories

import (
	"context"
	"fmt"

	"text.io/ent"
	"text.io/ent/link"
	"text.io/internal/entities/models"
)

type LinkRepository struct {
	client *ent.Client
}

func NewPostgresLinkRepository(client *ent.Client) *LinkRepository {
	return &LinkRepository{
		client: client,
	}
}

func (r *LinkRepository) CheckIfLinkExists(fingerprint string) (bool, error) {
	exists, err := r.client.Link.Query().Where(link.Fingerprint(fingerprint)).Exist(context.Background())
	if err != nil {
		return false, fmt.Errorf("failed checking if link exists: %w", err)
	}

	return exists, nil
}

func (r *LinkRepository) GetLink(fingerprint string) (*ent.Link, error) {
	link, err := r.client.Link.
		Query().
		Where(link.Fingerprint(fingerprint)).
		Only(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed getting link: %w", err)
	}
	return link, nil
}

func (r *LinkRepository) CreateLink(shortUrl string, link models.CreateLink) error {
	_, err := r.client.Link.
		Create().
		SetFingerprint(shortUrl).
		SetURL(link.Url).
		Save(context.Background())
	if err != nil {
		return fmt.Errorf("failed creating link: %w", err)
	}
	return nil
}

func (r *LinkRepository) ListLinks() ([]*ent.Link, error) {
	links, err := r.client.Link.
		Query().
		All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed quering the links: %w", err)
	}
	return links, nil
}
