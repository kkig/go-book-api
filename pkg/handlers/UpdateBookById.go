package handlers

import (
	"bookApi/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateBookById(ctx *gin.Context) {
	var input database.UpdateBookInput

	// Validate input
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Query by id before update
	book, err := database.FindOneyId(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBook, err := book.UpdateOne(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updatedBook})
}