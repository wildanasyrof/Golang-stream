package repository

import (
	"github.com/wildanasyrof/golang-stream/internal/entity"
	"gorm.io/gorm"
)

type EpisodeRepository interface {
	Create(episode *entity.Episode) error
	FindByID(id uint) (*entity.Episode, error)
	FindByAnimeID(animeID uint) ([]entity.Episode, error)
	Update(episode *entity.Episode) error
	Delete(id uint) error
}

type episodeRepository struct {
	db *gorm.DB
}

func NewEpisodeRepository(db *gorm.DB) EpisodeRepository {
	return &episodeRepository{db}
}

func (r *episodeRepository) Create(episode *entity.Episode) error {
	return r.db.Create(episode).Error
}

func (r *episodeRepository) FindByID(id uint) (*entity.Episode, error) {
	var episode entity.Episode
	err := r.db.First(&episode, id).Error
	if err != nil {
		return nil, err
	}
	return &episode, nil
}

func (r *episodeRepository) FindByAnimeID(animeID uint) ([]entity.Episode, error) {
	var episodes []entity.Episode
	err := r.db.Where("anime_id = ?", animeID).Find(&episodes).Error
	if err != nil {
		return nil, err
	}
	return episodes, nil
}

func (r *episodeRepository) Update(episode *entity.Episode) error {
	return r.db.Save(episode).Error
}

func (r *episodeRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Episode{}, id).Error
}
