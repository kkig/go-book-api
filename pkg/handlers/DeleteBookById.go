package handlers

import (
	"bookApi/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteBookById(ctx *gin.Context) {
	book, err := database.FindOneyId(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err = book.DeleteOne()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Deleted data."})
}