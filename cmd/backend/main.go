package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/leoanicio/deck_handler/pkg/api"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/create", api.CreateDeck)
	router.GET("/get/:deck_id", api.GetDeck)
	router.POST("/draw", api.DrawCard)

	return router
}

func main() {
	router := setupRouter()
	router.Run("localhost:8080")
}
