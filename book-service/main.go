package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	httpPort := ":" + os.Getenv("HTTP_PORT")

	route.GET("/test", func(ctx *gin.Context) {
		data := map[string]interface{}{
			"lang": "Hello语言",
			"tag": "<br>",
		}
		ctx.JSON(http.StatusOK, data)
		// ctx.JSON(http.StatusOK, gin.H{"data": "Hit!"})
	})

	route.Run(httpPort)
}