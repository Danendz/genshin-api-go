package main

import (
	"github.com/Danendz/genshin-api-go/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	server := gin.Default()

	if err := server.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal("Failed to set trusted proxies: ", err)
	}

	server.GET("/health-check", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	})

	v1 := server.Group("/v1")

	v1.GET("/characters", handlers.GetCharacters)

	v1.GET("/characters/:id", handlers.GetCharacter)

	if err := server.Run("localhost:8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
