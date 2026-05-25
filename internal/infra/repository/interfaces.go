package repository

import (
	"context"

	"github.com/devlucas-java/go-grpc/internal/domain"
)

type CourseRepository interface {
	FindByID(ctx context.Context, id string) (*domain.Course, error)
	FindAll(ctx context.Context) ([]*domain.Course, error)
	Create(course *domain.Course) (*domain.Course, error)
}

type CategoryRepository interface {
	FindAll(ctx context.Context, limit int) ([]*domain.Category, error)
	FindByID(ctx context.Context, id string) (*domain.Category, error)
	Create(category *domain.Category) (*domain.Category, error)
}
