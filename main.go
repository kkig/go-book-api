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

	v1 := route.Group("/v1")
	{
		v1.GET("/books", handlers.FindBooks)
		v1.GET("/books/:id", handlers.FindBookById)

		v1.POST("/books", handlers.AddBook)

		v1.PATCH("/books/:id", handlers.UpdateBookById)

		v1.DELETE("/books/:id", handlers.DeleteBookById)
	}

	route.Run()
}