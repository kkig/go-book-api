package database

import "gorm.io/gorm"

// GORM defined a gorm.Model struct, which includes fields:
// ID, CreatedAt, UpdatedAt, DeletedAt
type Book struct {
	gorm.Model
	Title	string	`json:"title" gorm:"size:255;not null" `
	Author	string	`json:"author" gorm:"size:255;not null"`
	Desc	string	`json:"desc" gorm:"size:255"`
}
type NewBookInput struct {
	Title		string	`json:"title" binding:"required"`
	Author		string	`json:"author" binding:"required"`
}