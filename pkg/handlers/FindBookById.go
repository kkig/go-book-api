package handlers

import (
	"bookApi/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindBookById(ctx *gin.Context) {
	book, err := database.FindBookById(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}