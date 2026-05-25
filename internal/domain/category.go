package domain

import "github.com/google/uuid"

type Category struct {
	ID          uuid.UUID
	Name        string
	Description string
}

func NewCategory(name string, description string) *Category {
	return &Category{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
	}
}
