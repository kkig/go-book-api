// You shouldn’t export the whole list outside this package
// to keep the module loosely coupled as microservice.

package database

func FindBookById(id string) (Book, error) {
	var book Book

	err := db.Where("ID = ?", id).First(&book).Error
	println(err)
	if err != nil {
		return Book{}, err
	}

	return book, nil
}

func FindBooks() ([]Book, error) {
	var books []Book

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

func (book *Book) UpdateBook(input *UpdateBookInput) (*Book, error) {
	err := db.Model(&book).Updates(&input).Error
	if err != nil {
		return &Book{}, err
	}

	return book, nil
}

func (book *Book) DeleteOne() (error) {
	err := db.Delete(&book).Error
	if err != nil {
		return err
	}

	return nil
}