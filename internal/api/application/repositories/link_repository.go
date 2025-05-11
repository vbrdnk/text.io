package repositories

import (
	"text.io/ent"
	"text.io/internal/entities/models"
)

type LinkRepository interface {
	CheckIfLinkExists(fingerprint string) (bool, error)
	GetLink(fingerprint string) (*ent.Link, error)
	CreateLink(shortUrl string, link models.CreateLink) error
	ListLinks() ([]*ent.Link, error)
}
