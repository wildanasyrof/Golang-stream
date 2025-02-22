package service

import (
	"errors"

	"github.com/wildanasyrof/golang-stream/internal/dto"
	"github.com/wildanasyrof/golang-stream/internal/entity"
	"github.com/wildanasyrof/golang-stream/internal/repository"
)

type AnimeService interface {
	CreateAnime(req dto.CreateAnimeRequest) (*entity.Anime, error)
	GetAllAnime(limit, page int, filters map[string]string) ([]dto.AnimeResponse, int64, error)
	GetAnimeByID(id uint) (*entity.Anime, error)
	UpdateAnime(id uint, req dto.UpdateAnimeRequest) (*entity.Anime, error)
	DeleteAnime(id uint) (*entity.Anime, error)
}

type animeService struct {
	animeRepo    repository.AnimeRepository
	categoryRepo repository.CategoryRepository
}

func NewAnimeService(animeRepo repository.AnimeRepository, categoryRepo repository.CategoryRepository) AnimeService {
	return &animeService{animeRepo, categoryRepo}
}

// Create a new anime with categories
func (s *animeService) CreateAnime(req dto.CreateAnimeRequest) (*entity.Anime, error) {
	// Validate category existence
	var categories []entity.Category
	for _, categoryName := range req.Categories {
		category, err := s.categoryRepo.FindByName(categoryName)
		if err != nil {
			return nil, errors.New("category " + categoryName + " not found")
		}
		categories = append(categories, *category)
	}

	anime := &entity.Anime{
		Title:       req.Title,
		AltTitles:   req.AltTitles,
		Chapters:    req.Chapters,
		Studio:      req.Studio,
		Year:        req.Year,
		Rating:      req.Rating,
		Synopsis:    req.Synopsis,
		ImageSource: req.ImageSource,
		Categories:  categories,
	}

	if err := s.animeRepo.Create(anime); err != nil {
		return nil, err
	}

	return anime, nil
}

// Get all anime with categories
func (s *animeService) GetAllAnime(limit, page int, filters map[string]string) ([]dto.AnimeResponse, int64, error) {
	// Set default limit if not provided
	if limit <= 0 {
		limit = 10
	}

	// Calculate offset
	offset := (page - 1) * limit

	animes, total, err := s.animeRepo.GetAllAnime(limit, offset, filters)
	if err != nil {
		return nil, 0, err
	}

	var animeList []dto.AnimeResponse
	for _, anime := range animes {
		animeList = append(animeList, dto.AnimeResponse{
			ID:           anime.ID,
			Title:        anime.Title,
			AltTitles:    anime.AltTitles,
			Chapters:     anime.Chapters,
			Studio:       anime.Studio,
			Year:         anime.Year,
			Rating:       anime.Rating,
			Synopsis:     anime.Synopsis,
			ImageSource:  anime.ImageSource,
			EpisodeCount: len(anime.Episodes), // Use len() for episode count
			Categories:   mapCategoriesToDTO(anime.Categories),
		})
	}

	return animeList, total, nil
}

// Helper function to convert categories to DTO
func mapCategoriesToDTO(categories []entity.Category) []dto.CategoryResponse {
	var categoryDTOs []dto.CategoryResponse
	for _, category := range categories {
		categoryDTOs = append(categoryDTOs, dto.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}
	return categoryDTOs
}

// Get anime by ID
func (s *animeService) GetAnimeByID(id uint) (*entity.Anime, error) {
	return s.animeRepo.FindByID(id)
}

// Update anime details
func (s *animeService) UpdateAnime(id uint, req dto.UpdateAnimeRequest) (*entity.Anime, error) {
	anime, err := s.animeRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("anime not found")
	}

	// Update fields if provided
	if req.Title != "" {
		anime.Title = req.Title
	}
	if req.AltTitles != "" {
		anime.AltTitles = req.AltTitles
	}
	if req.Chapters != "" {
		anime.Chapters = req.Chapters
	}
	if req.Studio != "" {
		anime.Studio = req.Studio
	}
	if req.Year != "" {
		anime.Year = req.Year
	}
	if req.Rating != 0 {
		anime.Rating = req.Rating
	}
	if req.Synopsis != "" {
		anime.Synopsis = req.Synopsis
	}
	if req.ImageSource != "" {
		anime.ImageSource = req.ImageSource
	}

	// Handle categories update
	if len(req.Categories) > 0 {
		var categories []entity.Category
		for _, categoryName := range req.Categories {
			category, err := s.categoryRepo.FindByName(categoryName)
			if err != nil {
				return nil, errors.New("category " + categoryName + " not found")
			}
			categories = append(categories, *category)
		}
		anime.Categories = categories
	}

	if err := s.animeRepo.Update(anime); err != nil {
		return nil, err
	}

	return anime, nil
}

// Delete anime by ID
func (s *animeService) DeleteAnime(id uint) (*entity.Anime, error) {
	anime, err := s.animeRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("anime not found")
	}

	if err := s.animeRepo.Delete(id); err != nil {
		return nil, err
	}

	return anime, nil
}
