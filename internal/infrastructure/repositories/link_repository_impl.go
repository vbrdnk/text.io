package repositories

import (
	"database/sql"

	"text.io/internal/entities/models"
)

type LinkRepository struct {
	db *sql.DB
}

func NewPostgresLinkRepository(db *sql.DB) *LinkRepository {
	return &LinkRepository{
		db: db,
	}
}

func (r *LinkRepository) GetLink(fingerprint string) (models.Link, error) {
	var link models.Link
	err := r.db.QueryRow("SELECT fingerprint FROM links WHERE fingerprint = $1", fingerprint).
		Scan(&link.Fingerprint)

	return link, err
}

func (r *LinkRepository) CreateLink(shortUrl string, link models.CreateLink) error {
	_, err := r.db.Exec("INSERT INTO links (fingerprint, url) VALUES ($1, $2)", shortUrl, link.Url)
	return err
}

func (r *LinkRepository) ListLinks() ([]models.Link, error) {
	rows, err := r.db.Query("SELECT fingerprint, url, created_at FROM links")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Parse rows into links
	var links []models.Link
	for rows.Next() {
		var link models.Link
		if err := rows.Scan(&link.Fingerprint, &link.Url, &link.CreatedAt); err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	return links, nil
}
