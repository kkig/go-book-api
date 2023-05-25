package main

import (
	"bookApi/pkg/handlers"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

func main() {
	route := gin.Default()
	// err := godotenv.Load(".env.local")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	httpPort := ":" + os.Getenv("HTTP_PORT")

	fmt.Printf("HTTP_PORT -> %v\n", httpPort)

	route.GET("/test", func(ctx *gin.Context) {
		data := map[string]interface{}{
			"lang": "Hello语言",
			"tag": "<br>",
		}
		ctx.JSON(http.StatusOK, data)
		// ctx.JSON(http.StatusOK, gin.H{"data": "Hit!"})
	})

	v1 := route.Group("/v1")
	{
		v1.GET("/books", handlers.GetAllBooks)
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "pong"})
		})
	}

	route.Run(httpPort)
}