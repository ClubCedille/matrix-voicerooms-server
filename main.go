package main

import (
	"net/http"

	"github.com/ClubCedille/matrix-voicerooms-server/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Matrix Voicerooms Web Server
// @version 2.0
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	router := gin.Default()
	router.GET("/health", getHealth)
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run("localhost:8080")
}

// @Summary Get Health
// @Router /health [get]
func getHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}
