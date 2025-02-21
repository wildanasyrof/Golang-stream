package repository

import (
	"github.com/wildanasyrof/golang-stream/internal/entity"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *entity.Category) error
	GetAll() ([]entity.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

// create implements CategoryRepository.
func (c *categoryRepository) Create(category *entity.Category) error {
	return c.db.Create(category).Error
}

// GetAll implements CategoryRepository.
func (c *categoryRepository) GetAll() ([]entity.Category, error) {
	var categories []entity.Category
	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
