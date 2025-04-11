package entity

import "time"

type Post struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title" validate:"required,min=20"`
	Content     string    `gorm:"type:text" json:"content" validate:"required,min=200"`
	Category    string    `gorm:"not null" json:"category" validate:"required,min=3"`
	Status      string    `gorm:"not null;default:'draft'" json:"status"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}
