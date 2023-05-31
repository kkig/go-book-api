package database

import (
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

	db.AutoMigrate(&Book{})

	fmt.Println("Connected to server!")
}

func Teardown() {
	migrator := db.Migrator()
	migrator.DropTable(&Book{})
}