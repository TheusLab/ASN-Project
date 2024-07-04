package main

import (
	"github.com/TheusLab/ASN-Project/backend/handlers"
	"github.com/TheusLab/ASN-Project/backend/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Custom logger middleware
	router.Use(func(c *gin.Context) {
		utils.Log.Info().Str("path", c.Request.URL.Path).Msg("Request received")
		c.Next()
	})

	router.GET("/search", handlers.Search)

	utils.Log.Info().Msg("Server running on port 8080")
	router.Run(":8080")
}
