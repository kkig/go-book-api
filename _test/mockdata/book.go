package mockdata

import "bookApi/pkg/models"

var Books = []models.Book{
	{
		Id: 1,
		Title: "The quest to 'Go'",
		Author: "Gopher",
		Desc: "A book for Go",
	},
}