package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var (
	port    = os.Getenv("PORT")
	version = os.Getenv("VERSION")
)

func main() {
	if port == "" {
		log.Fatal().Msg("$PORT must be set")
	}

	if version == "" {
		log.Fatal().Msg("$VERSION is not set")
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Greeter Service!",
		})
	})

	router.GET("/greeter", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"greetings": fmt.Sprintf("Hello from version %s", version),
		})
	})

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server:")
	}
}
