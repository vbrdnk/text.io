package services

import (
	"net/url"

	"text.io/ent"
	"text.io/internal/api/application/repositories"
	entity_errors "text.io/internal/entities/errors"
	"text.io/internal/entities/models"

	"text.io/pkg"
)

type LinksService struct {
	repo repositories.LinkRepository
}

func NewLinksService(repo repositories.LinkRepository) *LinksService {
	return &LinksService{
		repo: repo,
	}
}

func (s *LinksService) CreateLink(link models.CreateLink) error {
	// Validate the link
	_, err := url.ParseRequestURI(link.Url)
	if err != nil {
		return entity_errors.ErrInvalidUrl
	}
	shortUrl := pkg.GenerateURLFingerprint(link.Url)

	exists, err := s.repo.CheckIfLinkExists(shortUrl)
	if err != nil || exists {
		return entity_errors.ErrLinkExists
	}

	return s.repo.CreateLink(shortUrl, link)
}

func (s *LinksService) GetLink(fingerprint string) (*ent.Link, error) {
	return s.repo.GetLink(fingerprint)
}

func (s *LinksService) ListLinks() ([]*ent.Link, error) {
	return s.repo.ListLinks()
}
