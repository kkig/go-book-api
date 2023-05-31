// You shouldn’t export the whole list outside this package
// to keep the module loosely coupled as microservice.

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

// GORM defined a gorm.Model struct, which includes fields:
// ID, CreatedAt, UpdatedAt, DeletedAt
// type Book struct {
// 	gorm.Model
// 	Title	string	`json:"title" gorm:"size:255;not null" `
// 	Author	string	`json:"author" gorm:"size:255;not null"`
// 	Desc	string	`json:"desc" gorm:"size:255"`
// }
// type NewBookInput struct {
// 	Title		string	`json:"title" binding:"required"`
// 	Author		string	`json:"author" binding:"required"`
// }

var db *gorm.DB
// var book models.Book

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
	// var book models.Book
	db.AutoMigrate(&Book{})

	fmt.Println("Connected to server!")
}

func Teardown() {
	migrator := db.Migrator()
	migrator.DropTable(&Book{})
}

func FindBookById(id string) (Book, error) {
	var book = Book{}

	err := db.Where("ID = ?", id).First(&book).Error
	println(err)
	if err != nil {
		return Book{}, err
	}

	return book, nil
}

func FindBooks() ([]Book, error) {
	var books = []Book{}

	err := db.Find(&books).Error
	if err != nil {
		return []Book{}, err
	}

	return books, nil
}

func (book *Book) CreateOne() (*Book, error) {
	err := db.Create(&book).Error
	if err != nil {
		return &Book{}, err
	}

	return book, nil
}