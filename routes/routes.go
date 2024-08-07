package routes

import (
	"playlist-music/controllers"
	"playlist-music/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/register", controllers.RegisterUser)
	r.POST("/api/login", controllers.LoginUser)

	auth := r.Group("/")
	auth.Use(middleware.Auth())
	{
		auth.POST("/api/music", controllers.CreateMusic)
		auth.GET("/api/music", controllers.GetAllMusic)
		auth.GET("/api/music/:id", controllers.GetMusicByID)
		auth.PUT("/api/music/:id", controllers.UpdateMusic)
		auth.DELETE("/api/music/:id", controllers.DeleteMusic)

		auth.POST("/api/playlists", controllers.CreatePlaylist)
		auth.GET("/api/playlists", controllers.GetAllPlaylists)
		auth.GET("/api/playlists/:playlistID", controllers.GetPlaylistDetail)
		auth.PUT("/api/playlists/:playlistID", controllers.UpdatePlaylist)
		auth.DELETE("/api/playlists/:playlistID", controllers.DeletePlaylist)
		auth.POST("/api/playlists/:playlistID/music", controllers.AddMusicToPlaylist)
		auth.DELETE("/api/playlists/:playlistID/music/:musicID", controllers.RemoveMusicFromPlaylist)
	}

	return r
}
