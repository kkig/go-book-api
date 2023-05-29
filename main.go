package main

import (
	"fmt"
	"os"

	"bookApi/pkg/database"
	"bookApi/pkg/handlers"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

func main() {
	database.Connect()
	serveApplication()
}

func serveApplication() {
	route := gin.Default()
	// err := godotenv.Load(".env.local")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	fmt.Printf("HTTP_PORT -> %v\n", httpPort)

	// route.GET("/test", func(ctx *gin.Context) {
	// 	data := map[string]interface{}{
	// 		"lang": "Hello语言",
	// 		"tag": "<br>",
	// 	}
	// 	ctx.JSON(http.StatusOK, data)
	// 	// ctx.JSON(http.StatusOK, gin.H{"data": "Hit!"})
	// })

	v1 := route.Group("/v1")
	{
		v1.GET("/books", handlers.GetAllBooks)
		v1.POST("/books", handlers.AddBook)

		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "pong"})
		})
	}

	route.Run()
}