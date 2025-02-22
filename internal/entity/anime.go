package entity

import (
	"time"
)

type Anime struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"not null" json:"title"`
	AltTitles   string     `json:"alt_titles"`
	Chapters    string     `json:"chapters"`
	Studio      string     `json:"studio"`
	Year        string     `json:"year"`
	Rating      float64    `json:"rating"`
	Synopsis    string     `json:"synopsis"`
	ImageSource string     `json:"image_source"`
	Categories  []Category `gorm:"many2many:anime_categories;" json:"categories"`
	Episodes    []Episode  `json:"episodes"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
