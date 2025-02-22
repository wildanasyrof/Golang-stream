package service

import (
	"errors"

	"github.com/wildanasyrof/golang-stream/internal/dto"
	"github.com/wildanasyrof/golang-stream/internal/repository"
)

type FavoriteService interface {
	AddFavorite(userID, animeID uint) error
	RemoveFavorite(userID, animeID uint) error
	GetUserFavorites(userID uint, limit, offset int) ([]dto.AnimeResponse, int64, error)
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
func (s *favoriteService) GetUserFavorites(userID uint, limit, offset int) ([]dto.AnimeResponse, int64, error) {
	favorites, total, err := s.favoriteRepo.GetUserFavorites(userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	var animeList []dto.AnimeResponse
	for _, fav := range favorites {
		animeList = append(animeList, dto.AnimeResponse{
			ID:           fav.Anime.ID,
			Title:        fav.Anime.Title,
			AltTitles:    fav.Anime.AltTitles,
			Chapters:     fav.Anime.Chapters,
			Studio:       fav.Anime.Studio,
			Year:         fav.Anime.Year,
			Rating:       fav.Anime.Rating,
			Synopsis:     fav.Anime.Synopsis,
			ImageSource:  fav.Anime.ImageSource,
			EpisodeCount: len(fav.Anime.Episodes), // Use len() for episode count
			Categories:   mapCategoriesToDTO(fav.Anime.Categories),
		})
	}

	return animeList, total, nil
}
