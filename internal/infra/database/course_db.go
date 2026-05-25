package database

import (
	"context"
	"database/sql"

	"github.com/devlucas-java/go-grpc/internal/domain"
	"github.com/devlucas-java/go-grpc/internal/infra/repository"
)

type CourseDB struct {
	DB *sql.DB
}

func NewCourseDB(db *sql.DB) repository.CourseRepository {
	return &CourseDB{DB: db}
}

func (c *CourseDB) FindByID(ctx context.Context, id string) (*domain.Course, error) {
	var course domain.Course
	var category domain.Category

	err := c.DB.QueryRowContext(ctx, `
		SELECT c.id, c.name, c.description,
		       ca.id, ca.name, ca.description
		FROM courses c
		JOIN categories ca ON ca.id = c.category_id
		WHERE c.id = $1
	`, id).Scan(
		&course.ID,
		&course.Name,
		&course.Description,
		&category.ID,
		&category.Name,
		&category.Description,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	course.Category = &category
	return &course, nil
}

func (c *CourseDB) Create(course *domain.Course) (*domain.Course, error) {
	_, err := c.DB.Exec(
		`INSERT INTO courses (id, name, description, category_id)
		 VALUES ($1, $2, $3, $4)`,
		course.ID,
		course.Name,
		course.Description,
		course.Category.ID,
	)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (c *CourseDB) FindAll(ctx context.Context) ([]*domain.Course, error) {
	rows, err := c.DB.QueryContext(ctx, `
		SELECT c.id, c.name, c.description,
		       ca.id, ca.name, ca.description
		FROM courses c
		JOIN categories ca ON ca.id = c.category_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*domain.Course

	for rows.Next() {
		var course domain.Course
		var category domain.Category

		if err := rows.Scan(
			&course.ID,
			&course.Name,
			&course.Description,
			&category.ID,
			&category.Name,
			&category.Description,
		); err != nil {
			return nil, err
		}

		course.Category = &category
		courses = append(courses, &course)
	}

	return courses, rows.Err()
}
