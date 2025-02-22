package repository

import (
	"github.com/wildanasyrof/golang-stream/internal/entity"
	"gorm.io/gorm"
)

type AnimeRepository interface {
	Create(anime *entity.Anime) error
	GetAllAnime(limit, offset int, filters map[string]string) ([]entity.Anime, int64, error)
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

func (r *animeRepository) GetAllAnime(limit, offset int, filters map[string]string) ([]entity.Anime, int64, error) {
	var animes []entity.Anime
	var total int64

	query := r.db.Model(&entity.Anime{}).Preload("Categories").Preload("Episodes")

	// Apply Filters
	if title, ok := filters["title"]; ok {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if year, ok := filters["year"]; ok {
		query = query.Where("year = ?", year)
	}
	if studio, ok := filters["studio"]; ok {
		query = query.Where("studio LIKE ?", "%"+studio+"%")
	}
	if category, ok := filters["category"]; ok {
		query = query.Joins("JOIN anime_categories ON anime_categories.anime_id = animes.id").
			Joins("JOIN categories ON anime_categories.category_id = categories.id").
			Where("categories.name = ?", category)
	}

	// Count total results before pagination
	query.Count(&total)

	// Apply pagination
	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	err := query.Find(&animes).Error
	if err != nil {
		return nil, 0, err
	}

	return animes, total, nil
}

// Find anime by ID with categories
func (r *animeRepository) FindByID(id uint) (*entity.Anime, error) {
	var anime entity.Anime
	if err := r.db.Preload("Categories").Preload("Episodes").First(&anime, id).Error; err != nil {
		return nil, err
	}
	return &anime, nil
}

// Update anime
func (r *animeRepository) Update(anime *entity.Anime) error {
	return r.db.Save(anime).Error
}

// Delete anime by ID
func (r *animeRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Anime{}, id).Error
}
