package repository

import (
	"github.com/wildanasyrof/golang-stream/internal/entity"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *entity.Category) error
	GetAll() ([]entity.Category, error)
	FindByID(id uint) (*entity.Category, error)
	FindByName(name string) (*entity.Category, error)
	Destroy(id uint) error
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

// FindByID implements CategoryRepository.
func (c *categoryRepository) FindByID(id uint) (*entity.Category, error) {
	var category entity.Category
	if err := c.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) FindByName(name string) (*entity.Category, error) {
	var category entity.Category
	if err := r.db.Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// Destroy implements CategoryRepository.
func (c *categoryRepository) Destroy(id uint) error {
	return c.db.Delete(&entity.Category{}, id).Error
}
