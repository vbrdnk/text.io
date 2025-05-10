package repositories

import "text.io/internal/entities/models"

type LinkRepository interface {
	GetLink(fingerprint string) (models.Link, error)
	CreateLink(link models.Link) error
	ListLinks() ([]models.Link, error)
}
