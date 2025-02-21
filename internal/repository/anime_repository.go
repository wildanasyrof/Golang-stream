package repository

import (
	"errors"

	"github.com/wildanasyrof/golang-stream/internal/entity"
	"gorm.io/gorm"
)

type AnimeRepository interface {
	Create(anime *entity.Anime) error
	FindAll() ([]entity.Anime, error)
	FindByID(id uint) (*entity.Anime, error)
	Update(anime *entity.Anime) error
	Delete(id uint) error
}

type animeRepository struct {
	db *gorm.DB
}

func NewAnimeRepository(db *gorm.DB) AnimeRepository {
	return &animeRepository{db}
}

// Create anime with categories
func (r *animeRepository) Create(anime *entity.Anime) error {
	return r.db.Create(anime).Error
}

// Get all anime with categories
func (r *animeRepository) FindAll() ([]entity.Anime, error) {
	var animes []entity.Anime
	err := r.db.Preload("Categories").Find(&animes).Error
	return animes, err
}

// Find anime by ID with categories
func (r *animeRepository) FindByID(id uint) (*entity.Anime, error) {
	var anime entity.Anime
	err := r.db.Preload("Categories").First(&anime, id).Error
	if err != nil {
		return nil, errors.New("anime not found")
	}
	return &anime, nil
}

// Update anime
func (r *animeRepository) Update(anime *entity.Anime) error {
	return r.db.Save(anime).Error
}

// Delete anime by ID
func (r *animeRepository) Delete(id uint) error {
	result := r.db.Delete(&entity.Anime{}, id)
	if result.RowsAffected == 0 {
		return errors.New("anime not found")
	}
	return result.Error
}
