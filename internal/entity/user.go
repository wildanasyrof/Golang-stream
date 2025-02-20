package entity

import "time"

type User struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Username   string    `gorm:"unique;not null" json:"username"`
	Email      string    `gorm:"unique;not null" json:"email"`
	Password   string    `gorm:"not null" json:"-"` // Disembunyikan saat response
	Role       string    `gorm:"type:VARCHAR(10);default:'USER'" json:"role"`
	ProfileImg string    `gorm:"default:''" json:"profile_img"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
