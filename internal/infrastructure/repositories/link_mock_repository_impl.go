// internal/database/mock_repository.go
package repositories

import (
	"database/sql"
	"errors"

	"text.io/internal/entities/models"
)

// MockLinkRepository implements LinkRepository for testing
type MockLinkRepository struct {
	links map[string]models.Link
}

// NewMockLinkRepository creates a new mock repository
func NewMockLinkRepository() *MockLinkRepository {
	return &MockLinkRepository{
		links: make(map[string]models.Link),
	}
}

func (r *MockLinkRepository) GetLink(fingerprint string) (models.Link, error) {
	link, exists := r.links[fingerprint]
	if !exists {
		return models.Link{}, sql.ErrNoRows
	}
	return link, nil
}

func (r *MockLinkRepository) CreateLink(link models.Link) error {
	if _, exists := r.links[link.Fingerprint]; exists {
		return errors.New("link already exists")
	}
	r.links[link.Fingerprint] = link
	return nil
}

func (r *MockLinkRepository) ListLinks() ([]models.Link, error) {
	links := make([]models.Link, 0, len(r.links))
	for _, link := range r.links {
		links = append(links, link)
	}
	return links, nil
}
