package service

import (
	"errors"

	"github.com/wildanasyrof/golang-stream/internal/dto"
	"github.com/wildanasyrof/golang-stream/internal/entity"
	"github.com/wildanasyrof/golang-stream/internal/repository"
)

type EpisodeService interface {
	CreateEpisode(animeId uint, req dto.CreateEpisodeRequest) (*entity.Episode, error)
	GetEpisodeByID(id uint) (*entity.Episode, error)
	GetEpisodesByAnimeID(animeID uint) ([]entity.Episode, error)
	UpdateEpisode(animeId uint, epsNumber int, req dto.UpdateEpisodeRequest) (*entity.Episode, error)
	DeleteEpisode(animeId uint, epsNumber int) (*entity.Episode, error)
}

type episodeService struct {
	episodeRepo repository.EpisodeRepository
	animeRepo   repository.AnimeRepository
}

func NewEpisodeService(episodeRepo repository.EpisodeRepository, animeRepo repository.AnimeRepository) EpisodeService {
	return &episodeService{episodeRepo, animeRepo}
}

func (s *episodeService) CreateEpisode(animeId uint, req dto.CreateEpisodeRequest) (*entity.Episode, error) {
	// Check if anime exists
	anime, err := s.animeRepo.FindByID(animeId)
	if err != nil || anime == nil {
		return nil, errors.New("anime not found")
	}

	episode := &entity.Episode{
		AnimeID:       animeId,
		EpisodeNumber: req.EpisodeNumber,
		Title:         req.Title,
		VideoURL:      req.VideoURL,
	}

	err = s.episodeRepo.Create(episode)
	if err != nil {
		return nil, err
	}

	return episode, nil
}

func (s *episodeService) GetEpisodeByID(id uint) (*entity.Episode, error) {
	return s.episodeRepo.FindByID(id)
}

func (s *episodeService) GetEpisodesByAnimeID(animeID uint) ([]entity.Episode, error) {
	return s.episodeRepo.FindByAnimeID(animeID)
}

func (s *episodeService) UpdateEpisode(animeId uint, epsNumber int, req dto.UpdateEpisodeRequest) (*entity.Episode, error) {
	// Check if anime exists
	anime, err := s.animeRepo.FindByID(animeId)
	if err != nil || anime == nil {
		return nil, errors.New("anime not found")
	}

	episode, err := s.episodeRepo.FindByEpsNumber(animeId, epsNumber)
	if err != nil {
		return nil, errors.New("episode not found")
	}

	if req.EpisodeNumber != 0 {
		episode.EpisodeNumber = req.EpisodeNumber
	}
	if req.Title != "" {
		episode.Title = req.Title
	}
	if req.VideoURL != "" {
		episode.VideoURL = req.VideoURL
	}

	err = s.episodeRepo.Update(episode)
	if err != nil {
		return nil, err
	}

	return episode, nil
}

func (s *episodeService) DeleteEpisode(animeId uint, epsNumber int) (*entity.Episode, error) {
	// Check if anime exists
	anime, err := s.animeRepo.FindByID(animeId)
	if err != nil || anime == nil {
		return nil, errors.New("anime not found")
	}

	episode, err := s.episodeRepo.FindByEpsNumber(animeId, epsNumber)
	if err != nil {
		return nil, errors.New("episode not found")
	}

	if err := s.episodeRepo.Delete(episode.ID); err != nil {
		return nil, err

	}

	return episode, nil
}
