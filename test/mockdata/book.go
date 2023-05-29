package mockdata

import "bookApi/pkg/database"

var Books = []database.Book{
	{
		Title: "The quest to 'Go'",
		Author: "Gopher",
		Desc: "A book for Go",
	},
}