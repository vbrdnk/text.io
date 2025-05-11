package repositories

import "text.io/internal/entities/models"

type LinkRepository interface {
	GetLink(fingerprint string) (models.Link, error)
	CreateLink(shortUrl string, link models.CreateLink) error
	ListLinks() ([]models.Link, error)
}
