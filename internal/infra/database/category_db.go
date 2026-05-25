package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/devlucas-java/go-grpc/internal/domain"
	"github.com/devlucas-java/go-grpc/internal/infra/repository"
)

type CategoryDB struct {
	DB *sql.DB
}

func NewCategoryDB(db *sql.DB) repository.CategoryRepository {
	return &CategoryDB{DB: db}
}

func (c *CategoryDB) Create(category *domain.Category) (*domain.Category, error) {
	log.Printf("[DB] Creating category: %s", category.ID)

	_, err := c.DB.Exec(
		`INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)`,
		category.ID,
		category.Name,
		category.Description,
	)
	if err != nil {
		log.Printf("[DB][ERROR] Exec failed: %v", err)
		return nil, err
	}

	return category, nil
}

func (c *CategoryDB) FindByID(ctx context.Context, id string) (*domain.Category, error) {
	var category domain.Category

	err := c.DB.QueryRowContext(
		ctx,
		`SELECT id, name, description FROM categories WHERE id = $1`,
		id,
	).Scan(&category.ID, &category.Name, &category.Description)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (c *CategoryDB) FindAll(ctx context.Context, limit int) ([]*domain.Category, error) {
	rows, err := c.DB.QueryContext(
		ctx,
		`SELECT id, name, description FROM categories LIMIT $1`,
		limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*domain.Category

	for rows.Next() {
		var category domain.Category

		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}

	return categories, rows.Err()
}
