package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCharacters(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Characters",
	})
}

func GetCharacter(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Character id: " + id,
	})
}
