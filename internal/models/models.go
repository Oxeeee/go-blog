package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	Name             string         `gorm:"size:100;not null" json:"name"`
	Email            string         `gorm:"size:255;unique;not null" json:"email"`
	Password         string         `gorm:"size:255;not null" json:"password"`
	IsVerified       bool           `gorm:"default:false" json:"is_verified"`
	VerificationCode string         `gorm:"size:255" json:"verification_code"`
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

type Post struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null;index" json:"user_id"`
	Title     string         `gorm:"size:255;not null" json:"title"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Like struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;index"`
	PostID    uint      `gorm:"not null;index"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
