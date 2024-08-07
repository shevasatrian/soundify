package main

import (
	"playlist-music/database"
	"playlist-music/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Connect()

	r := routes.SetupRouter()

	r.Run(":8080")
}
