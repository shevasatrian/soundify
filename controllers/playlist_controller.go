package controllers

import (
	"net/http"
	"playlist-music/database"
	"playlist-music/models"

	"github.com/gin-gonic/gin"
)

// CreatePlaylist creates a new playlist
func CreatePlaylist(c *gin.Context) {
	var playlist models.Playlist
	if err := c.ShouldBindJSON(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil user_id dari context yang di-set oleh middleware
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	playlist.UserID = userID.(int)

	if err := database.DB.Create(&playlist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Preload user data
	if err := database.DB.Preload("User").First(&playlist, playlist.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, playlist)
}

// GetAllPlaylists retrieves all playlists
func GetAllPlaylists(c *gin.Context) {
	var playlists []models.Playlist
	if err := database.DB.Preload("User").Omit("Musics").Find(&playlists).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, playlists)
}

// GetPlaylistDetail retrieves a playlist by its ID, including its music
func GetPlaylistDetail(c *gin.Context) {
	playlistID := c.Param("playlistID")
	var playlist models.Playlist
	if err := database.DB.Preload("Musics").Preload("User").First(&playlist, "id = ?", playlistID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Playlist not found"})
		return
	}

	c.JSON(http.StatusOK, playlist)
}

// UpdatePlaylist updates a playlist
func UpdatePlaylist(c *gin.Context) {
	playlistID := c.Param("playlistID")
	var playlist models.Playlist
	if err := database.DB.Preload("Musics").Preload("User").First(&playlist, "id = ?", playlistID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Playlist not found"})
		return
	}

	if err := c.ShouldBindJSON(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&playlist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, playlist)
}

// DeletePlaylist deletes a playlist
func DeletePlaylist(c *gin.Context) {
	playlistID := c.Param("playlistID")
	var playlist models.Playlist

	// Check if the playlist exists
	if err := database.DB.First(&playlist, "id = ?", playlistID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Playlist not found"})
		return
	}

	// Delete the playlist
	if err := database.DB.Delete(&playlist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete playlist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "playlist deleted"})
}

// AddMusicToPlaylist adds music to a playlist
func AddMusicToPlaylist(c *gin.Context) {
	var input struct {
		MusicID int `json:"music_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playlistID := c.Param("playlistID")

	var playlist models.Playlist
	if err := database.DB.Preload("Musics").First(&playlist, "id = ?", playlistID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Playlist not found"})
		return
	}

	var music models.Music
	if err := database.DB.First(&music, "id = ?", input.MusicID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Music not found"})
		return
	}

	// Add music to playlist
	playlist.Musics = append(playlist.Musics, music)
	if err := database.DB.Save(&playlist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add music to playlist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Music added to playlist"})
}

// RemoveMusicFromPlaylist removes music from a playlist
func RemoveMusicFromPlaylist(c *gin.Context) {
	playlistID := c.Param("playlistID")
	musicID := c.Param("musicID")

	var playlist models.Playlist
	if err := database.DB.Preload("Musics").First(&playlist, "id = ?", playlistID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Playlist not found"})
		return
	}

	var music models.Music
	if err := database.DB.First(&music, "id = ?", musicID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Music not found"})
		return
	}

	// Remove music from playlist using Association
	if err := database.DB.Model(&playlist).Association("Musics").Delete(&music); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove music from playlist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "music deleted from playlist"})
}
