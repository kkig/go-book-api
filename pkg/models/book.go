package models

import "gorm.io/gorm"

// GORM defined a gorm.Model struct, which includes fields:
// ID, CreatedAt, UpdatedAt, DeletedAt
type Book struct {
	gorm.Model
	Title	string	`gorm:"size:255;not null" json:"title"`
	Author	string	`gorm:"size:255;not null" json:"author"`
	Desc	string	`gorm:"size:255" json:"desc"`
}
type NewBookInput struct {
	Title		string	`json:"title" binding:"required"`
	Author		string	`json:"author" binding:"required"`
}