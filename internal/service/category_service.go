package service

import (
	"github.com/wildanasyrof/golang-stream/internal/dto"
	"github.com/wildanasyrof/golang-stream/internal/entity"
	"github.com/wildanasyrof/golang-stream/internal/repository"
)

type CategoryService interface {
	CreateCategory(req dto.CreateCategoryRequest) (*entity.Category, error)
	GetAllCategories() ([]entity.Category, error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo}
}

// CreateCategory implements CategoryService.
func (c *categoryService) CreateCategory(req dto.CreateCategoryRequest) (*entity.Category, error) {
	category := &entity.Category{
		Name: req.Name,
	}

	if err := c.categoryRepo.Create(category); err != nil {
		return nil, err
	}

	return category, nil
}

// GetAllCategories implements CategoryService.
func (c *categoryService) GetAllCategories() ([]entity.Category, error) {
	return c.categoryRepo.GetAll()
}
