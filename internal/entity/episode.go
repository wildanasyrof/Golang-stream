package entity

import "time"

type Episode struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	AnimeID       uint      `gorm:"not null" json:"anime_id"`
	Anime         Anime     `gorm:"foreignKey:AnimeID" json:"-"`
	EpisodeNumber int       `gorm:"unique;not null" json:"episode_number"`
	Title         string    `gorm:"not null" json:"title"`
	VideoURL      string    `gorm:"not null" json:"video_url"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}
