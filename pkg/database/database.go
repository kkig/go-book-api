// You shouldn’t export the whole list outside this package
// to keep the module loosely coupled as microservice.

package database

import (
	"bookApi/pkg/models"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Connect() {
	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: logger.Silent,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries: true,
			Colorful: false,
		},
	)

	dbName := os.Getenv("PGDATABASE")
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	pass := os.Getenv("PGPASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, dbName, user, pass)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic(err)
	}

	var book models.Book
	db.AutoMigrate(&book)

	fmt.Println("Connected to server!")
}