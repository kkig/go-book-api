package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddBook(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "hit!"})
}