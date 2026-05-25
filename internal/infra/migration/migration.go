package migration

import (
	"database/sql"
	"log"
)

type Migration struct {
	db *sql.DB
}

func NewMigration(db *sql.DB) *Migration {
	return &Migration{db: db}
}

func (m *Migration) Run() error {
	log.Println("[MIGRATION] Running auto migrations...")

	if err := m.createCategoryTable(); err != nil {
		return err
	}

	log.Println("[MIGRATION] All migrations completed successfully")
	return nil
}

func (m *Migration) createCategoryTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS categories (
		id TEXT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description VARCHAR(255)
	);
	`

	if _, err := m.db.Exec(query); err != nil {
		log.Printf("[MIGRATION][ERROR] Failed to create categories table: %v", err)
		return err
	}

	log.Println("[MIGRATION] Categories table created successfully")
	return nil
}
