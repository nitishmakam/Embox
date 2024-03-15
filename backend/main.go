package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nitishmakam/embox/database"
	"github.com/nitishmakam/embox/handlers"
)

func main() {
	database.InitDB()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/movies/:id", handlers.GetMovie)
	router.GET("/movies", handlers.GetMovies)
	router.POST("/movies", handlers.AddMovie)
	router.PUT("/movies/:id", handlers.UpdateMovie)
	router.DELETE("/movies/:id", handlers.DeleteMovie)

	router.Run(":8443")

	defer database.ShutDownDB()
}
