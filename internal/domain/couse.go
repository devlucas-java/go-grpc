package domain

import "github.com/google/uuid"

type Course struct {
	ID          uuid.UUID
	Name        string
	Description string
	Category    *Category
}
