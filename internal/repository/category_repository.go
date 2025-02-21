package repository

import (
	"github.com/wildanasyrof/golang-stream/internal/entity"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *entity.Category) error
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
