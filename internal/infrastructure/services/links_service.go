package services

import (
	"errors"
	"fmt"
	"net/url"

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

	existingLink, err := s.repo.GetLink(shortUrl)
	fmt.Print(err)
	// TODO: revisit this logic and handle errors properly when db is empty
	if err != nil && existingLink.Fingerprint != "" {
		return entity_errors.ErrLinkExists
	}

	return s.repo.CreateLink(shortUrl, link)
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
