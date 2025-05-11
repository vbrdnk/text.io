package repositories

import (
	"database/sql"

	"text.io/internal/entities/models"
)

type CollectionRepository struct {
	db *sql.DB
}

func NewPostgresCollectionRepository(db *sql.DB) *CollectionRepository {
	return &CollectionRepository{
		db: db,
	}
}

func (r *CollectionRepository) GetCollection(fingerprint string) (models.Collection, error) {
	var collection models.Collection
	err := r.db.QueryRow("SELECT fingerprint FROM collections WHERE fingerprint = $1", fingerprint).
		Scan(&collection.Fingerprint)

	return collection, err
}

func (r *CollectionRepository) CreateCollection(shortUrl string, collection models.CreateCollection) error {
	_, err := r.db.Exec("INSERT INTO collections (fingerprint, title) VALUES ($1, $2)", shortUrl, collection.Title)
	return err
}

func (r *CollectionRepository) ListCollections() ([]models.Collection, error) {
	rows, err := r.db.Query("SELECT fingerprint, title, created_at FROM collections")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Parse rows into collections
	var collections []models.Collection
	for rows.Next() {
		var collection models.Collection
		if err := rows.Scan(&collection.Fingerprint, &collection.Title, &collection.CreatedAt); err != nil {
			return nil, err
		}

		collections = append(collections, collection)
	}

	return collections, nil
}
