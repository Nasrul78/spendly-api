package service

import (
	"context"

	"github.com/nasrul78/spendly-api/internal/domain"
	"github.com/nasrul78/spendly-api/internal/repository"
)

type CategoryService struct {
	categoryRepo *repository.CategoryRepository
}

func NewCategoryService(categoryRepo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (s *CategoryService) Create(ctx context.Context, userID string, req *domain.CreateCategoryRequest) (*domain.CategoryResponse, error) {
	category, err := s.categoryRepo.Create(ctx, userID, req.Name)
	if err != nil {
		return nil, err
	}

	return &domain.CategoryResponse{
		ID:        category.ID.Bytes,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Time,
		UpdatedAt: category.UpdatedAt.Time,
	}, nil
}

func (s *CategoryService) GetAll(ctx context.Context, userID string) ([]domain.CategoryResponse, error) {
	categories, err := s.categoryRepo.GetAllByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]domain.CategoryResponse, len(categories))
	for i, category := range categories {
		res[i] = domain.CategoryResponse{
			ID:        category.ID.Bytes,
			Name:      category.Name,
			CreatedAt: category.CreatedAt.Time,
			UpdatedAt: category.UpdatedAt.Time,
		}
	}

	return res, nil
}

func (s *CategoryService) GetByID(ctx context.Context, userID, categoryID string) (*domain.CategoryResponse, error) {
	category, err := s.categoryRepo.GetByID(ctx, userID, categoryID)
	if err != nil {
		return nil, err
	}

	return &domain.CategoryResponse{
		ID:        category.ID.Bytes,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Time,
		UpdatedAt: category.UpdatedAt.Time,
	}, nil
}

func (s *CategoryService) Update(ctx context.Context, userID, categoryID string, req *domain.UpdateCategoryRequest) (*domain.CategoryResponse, error) {
	category, err := s.categoryRepo.Update(ctx, userID, categoryID, req.Name)
	if err != nil {
		return nil, err
	}

	return &domain.CategoryResponse{
		ID:        category.ID.Bytes,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Time,
		UpdatedAt: category.UpdatedAt.Time,
	}, nil
}

func (s *CategoryService) Delete(ctx context.Context, userID, categoryID string) error {
	if err := s.categoryRepo.Delete(ctx, userID, categoryID); err != nil {
		return err
	}

	return nil
}
