package services

import (
	"errors"
	"net/url"

	"text.io/internal/api/application/repositories"
	entity_errors "text.io/internal/entities/errors"
	"text.io/internal/entities/models"
)

type LinksService struct {
	repo repositories.LinkRepository
}

func NewService(repo repositories.LinkRepository) *LinksService {
	return &LinksService{
		repo: repo,
	}
}

func (s *LinksService) CreateLink(link models.Link) error {
	// Validate the link
	_, err := url.ParseRequestURI(link.Url)
	if err != nil {
		return entity_errors.ErrInvalidUrl
	}

	existingLink, err := s.repo.GetLink(link.Fingerprint)
	if err != nil && existingLink.Fingerprint != "" {
		return entity_errors.ErrLinkExists
	}

	return s.repo.CreateLink(link)
}

func (s *LinksService) GetLink(fingerprint string) (models.Link, error) {
	if fingerprint == "" {
		return models.Link{}, errors.New("link Fingerprint is required")
	}

	return s.repo.GetLink(fingerprint)
}

func (s *LinksService) ListLinks() ([]models.Link, error) {
	return s.repo.ListLinks()
}
