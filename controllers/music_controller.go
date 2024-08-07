package controllers

import (
	"net/http"
	"playlist-music/database"
	"playlist-music/models"

	"github.com/gin-gonic/gin"
)

// CreateMusic creates a new music record
func CreateMusic(c *gin.Context) {
	var music models.Music
	if err := c.ShouldBindJSON(&music); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&music).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, music)
}

// GetAllMusic retrieves all music records
func GetAllMusic(c *gin.Context) {
	var musics []models.Music
	if err := database.DB.Find(&musics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, musics)
}

// GetMusicByID retrieves a single music record by its ID
func GetMusicByID(c *gin.Context) {
	var music models.Music
	id := c.Param("id")
	if err := database.DB.First(&music, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Music not found"})
		return
	}
	c.JSON(http.StatusOK, music)
}

// UpdateMusic updates a music record
func UpdateMusic(c *gin.Context) {
	var music models.Music
	id := c.Param("id")
	if err := database.DB.First(&music, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Music not found"})
		return
	}

	if err := c.ShouldBindJSON(&music); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&music).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, music)
}

// DeleteMusic deletes a music record
func DeleteMusic(c *gin.Context) {
	var music models.Music
	id := c.Param("id")
	if err := database.DB.Delete(&music, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Music deleted"})
}
