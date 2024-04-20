package main

import (
	"github.com/Danendz/genshin-api-go/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	server := gin.Default()

	if err := server.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal("Failed to set trusted proxies: ", err)
	}

	v1 := server.Group("/v1")

	v1.GET("/characters", handlers.GetCharacters)

	v1.GET("/characters/:id", handlers.GetCharacter)

	if err := server.Run("localhost:8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
