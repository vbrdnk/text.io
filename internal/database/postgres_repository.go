package database

import (
	"database/sql"

	"text.io/internal/models"
)

type PostgresItemRepository struct {
	db *sql.DB
}

func NewPostgresItemRepository(db *sql.DB) *PostgresItemRepository {
	return &PostgresItemRepository{
		db: db,
	}
}

func (r *PostgresItemRepository) GetItem(id string) (models.Item, error) {
	var item models.Item
	err := r.db.QueryRow("SELECT id, name FROM items WHERE id = $1", id).
		Scan(&item.ID, &item.Name)

	return item, err
}

func (r *PostgresItemRepository) CreateItem(item models.Item) error {
	_, err := r.db.Exec("INSERT INTO items (id, name) VALUES ($1, $2)", item.ID, item.Name)
	return err
}

func (r *PostgresItemRepository) ListItems() ([]models.Item, error) {
	rows, err := r.db.Query("SELECT id, name, created_at FROM items")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Parse rows into items
	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.CreatedAt); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}
