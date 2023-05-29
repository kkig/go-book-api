package handlers

import (
	"bookApi/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddBook(ctx *gin.Context) {
	var input database.NewBookInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	book := database.Book{
		Title:		input.Title,
		Author:		input.Author,
		Desc:		input.Desc,
	}

	savedBook, err := book.CreateOne()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": savedBook})
}