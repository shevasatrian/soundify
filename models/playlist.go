package models

import "time"

type Playlist struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	UserID      int       `json:"user_id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	Musics      []Music   `gorm:"many2many:playlist_music;" json:"musics"`
	User        User
}

type PlaylistMusic struct {
	PlaylistID int       `gorm:"primaryKey" json:"playlist_id"`
	MusicID    int       `gorm:"primaryKey" json:"music_id"`
	AddedAt    time.Time `gorm:"autoCreateTime" json:"added_at"`
	Playlist   Playlist
	Music      Music
}
