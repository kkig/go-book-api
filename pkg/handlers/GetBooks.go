package handlers

import (
	"net/http"

	"bookApi/_test/mockdata"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"user": mockdata.Books})
}