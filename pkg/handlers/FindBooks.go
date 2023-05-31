package handlers

import (
	"net/http"

	"bookApi/pkg/database"

	"github.com/gin-gonic/gin"
)

func FindBooks(ctx *gin.Context) {
	books, err := database.FindBooks()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": books})
}