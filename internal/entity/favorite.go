package entity

import "time"

type Favorite struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	AnimeID   uint      `gorm:"not null" json:"anime_id"`
	Anime     Anime     `gorm:"foreignKey:AnimeID" json:"anime"` // Preload anime details
	CreatedAt time.Time `json:"created_at"`
}
