package entity

import "time"

// Article represents the structure of the articles table in the database
type Article struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"not null" json:"title" validate:"required,min=20"`
	Content     string     `gorm:"not null" json:"content" validate:"required,min=200"`
	Category    string     `gorm:"not null" json:"category" validate:"required,min=3"`
	Status      string     `gorm:"not null;default:'Draft'" json:"status"`
	CreatedDate time.Time  `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate time.Time  `gorm:"column:updated_date;autoUpdateTime" json:"updated_date"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"` // For soft delete
}
