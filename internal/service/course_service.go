package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/devlucas-java/go-grpc/internal/domain"
	"github.com/devlucas-java/go-grpc/internal/infra/repository"
)

type CourseService struct {
	repo     repository.CourseRepository
	category repository.CategoryRepository
}

func NewCourseService(
	repo repository.CourseRepository,
	category repository.CategoryRepository,
) *CourseService {
	return &CourseService{
		repo:     repo,
		category: category,
	}
}

func (s *CourseService) Create(ctx context.Context, name, description, categoryID string) (*domain.Course, error) {

	category, err := s.category.FindByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, fmt.Errorf("category not found")
	}

	course := &domain.Course{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Category:    category,
	}

	_, err = s.repo.Create(course)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (s *CourseService) FindByID(ctx context.Context, id string) (*domain.Course, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *CourseService) FindAll(ctx context.Context) ([]*domain.Course, error) {
	return s.repo.FindAll(ctx)
}
