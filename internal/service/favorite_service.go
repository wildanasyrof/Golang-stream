package service

import (
	"errors"

	"github.com/wildanasyrof/golang-stream/internal/entity"
	"github.com/wildanasyrof/golang-stream/internal/repository"
)

type FavoriteService interface {
	AddFavorite(userID, animeID uint) error
	RemoveFavorite(userID, animeID uint) error
	GetUserFavorites(userID uint, limit, offset int) ([]entity.Favorite, int64, error)
}

type favoriteService struct {
	favoriteRepo repository.FavoriteRepository
	animeRepo    repository.AnimeRepository
}

func NewFavoriteService(favoriteRepo repository.FavoriteRepository, animeRepo repository.AnimeRepository) FavoriteService {
	return &favoriteService{favoriteRepo: favoriteRepo, animeRepo: animeRepo}
}

// AddFavorite adds an anime to the user's favorites
func (s *favoriteService) AddFavorite(userID, animeID uint) error {
	// Check if anime exists
	_, err := s.animeRepo.FindByID(animeID)
	if err != nil {
		return errors.New("anime not found")
	}

	// Check if already favorited
	isFav, err := s.favoriteRepo.IsFavorited(userID, animeID)
	if err != nil {
		return err
	}
	if isFav {
		return errors.New("anime is already in favorites")
	}

	return s.favoriteRepo.AddFavorite(userID, animeID)
}

// RemoveFavorite removes an anime from the user's favorites
func (s *favoriteService) RemoveFavorite(userID, animeID uint) error {
	// Check if anime is favorited before deleting
	isFav, err := s.favoriteRepo.IsFavorited(userID, animeID)
	if err != nil {
		return err
	}
	if !isFav {
		return errors.New("anime is not in favorites")
	}

	return s.favoriteRepo.RemoveFavorite(userID, animeID)
}

// GetUserFavorites retrieves paginated favorite animes of a user
func (s *favoriteService) GetUserFavorites(userID uint, limit, offset int) ([]entity.Favorite, int64, error) {
	return s.favoriteRepo.GetUserFavorites(userID, limit, offset)
}
