package repository

import (
	"github.com/wildanasyrof/golang-stream/internal/entity"
	"gorm.io/gorm"
)

type FavoriteRepository interface {
	AddFavorite(userID, animeID uint) error
	RemoveFavorite(userID, animeID uint) error
	IsFavorited(userID, animeID uint) (bool, error)
	GetUserFavorites(userID uint, limit, offset int) ([]entity.Favorite, int64, error)
}

type favoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) FavoriteRepository {
	return &favoriteRepository{db}
}

// AddFavorite adds a new favorite entry
// AddFavorite adds an anime to the user's favorites
func (r *favoriteRepository) AddFavorite(userID, animeID uint) error {
	favorite := entity.Favorite{UserID: userID, AnimeID: animeID}
	return r.db.Create(&favorite).Error
}

// RemoveFavorite deletes a favorite entry
func (r *favoriteRepository) RemoveFavorite(userID, animeID uint) error {
	return r.db.Where("user_id = ? AND anime_id = ?", userID, animeID).Delete(&entity.Favorite{}).Error
}

// IsFavorited checks if an anime is already favorited by the user
func (r *favoriteRepository) IsFavorited(userID, animeID uint) (bool, error) {
	var count int64
	err := r.db.Model(&entity.Favorite{}).Where("user_id = ? AND anime_id = ?", userID, animeID).Count(&count).Error
	return count > 0, err
}

// GetUserFavorites retrieves all favorite animes of a user with pagination
func (r *favoriteRepository) GetUserFavorites(userID uint, limit, offset int) ([]entity.Favorite, int64, error) {
	var total int64

	query := r.db.Model(&entity.Favorite{}).Where("user_id = ?", userID).Count(&total)

	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	var results []entity.Favorite
	err := query.Preload("Anime.Categories").Preload("Anime.Episodes").Find(&results).Error
	if err != nil {
		return nil, 0, err
	}

	return results, total, nil
}
