package handlers

import (
	"net/http"

	"bookApi/pkg/database"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(ctx *gin.Context) {
	books, err := database.GetAllBooks()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": books})
}