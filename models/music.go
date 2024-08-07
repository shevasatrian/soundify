package models

import "time"

type Music struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Artist      string    `gorm:"not null" json:"artist"`
	Album       string    `json:"album"`
	Genre       string    `json:"genre"`
	Duration    int       `json:"duration"`
	ReleaseYear int       `json:"release_year"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
