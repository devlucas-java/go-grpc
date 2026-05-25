package service

import (
	"context"
	"io"

	"github.com/devlucas-java/go-grpc/internal/delivery/grpc/pb"
	"github.com/devlucas-java/go-grpc/internal/domain"
	"github.com/devlucas-java/go-grpc/internal/infra/repository"
)

type CategoryService struct {
	repo repository.CategoryRepository
	pb.UnimplementedCategoryServiceServer
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category := domain.NewCategory(in.Name, in.Description)

	created, err := s.repo.Create(category)
	if err != nil {
		return nil, err
	}
	response := &pb.Category{
		Id:          created.ID.String(),
		Name:        created.Name,
		Description: created.Description,
	}
	return &pb.CategoryResponse{Category: response}, nil
}

func (s *CategoryService) FindByID(ctx context.Context, in *pb.Search) (*pb.CategoryResponse, error) {
	category, err := s.repo.FindByID(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.Category{
		Id:          category.ID.String(),
		Name:        category.Name,
		Description: category.Description,
	}
	return &pb.CategoryResponse{Category: response}, nil
}
func (s *CategoryService) ListAll(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := s.repo.FindAll(ctx, 20)
	if err != nil {
		return nil, err
	}
	var categoryList []*pb.Category

	for _, c := range categories {

		var category pb.Category

		category.Id = c.ID.String()
		category.Name = c.Name
		category.Description = c.Description

		categoryList = append(categoryList, &category)
	}

	return &pb.CategoryList{Categories: categoryList}, nil
}

func (s *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	createdCategories := &pb.CategoryList{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(createdCategories)
		}
		if err != nil {
			return err
		}

		category, err := s.repo.Create(domain.NewCategory(
			req.Name, req.Description,
		))
		if err != nil {
			return err
		}

		createdCategories.Categories = append(createdCategories.Categories, &pb.Category{
			Id:          category.ID.String(),
			Name:        category.Name,
			Description: category.Description,
		})
	}
}

func (s *CategoryService) CreateCategoryStramBidirectional(stream pb.CategoryService_CreateCategoryStramBidirectionalServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		category, err := s.repo.Create(domain.NewCategory(req.Name, req.Description))
		if err != nil {
			return err
		}

		if err := stream.Send(&pb.Category{
			Id:          category.ID.String(),
			Name:        category.Name,
			Description: category.Description,
		}); err != nil {
			return err
		}
	}
}
