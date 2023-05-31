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

func (input *UpdateBookInput) UpdateBook(id string) (error) {
	var book Book
	err := db.Model(&book).Where("ID = ?", id).Updates(&input).Error
	if err != nil {
		return err
	}

	return nil
}

func (book *Book) DeleteOne() (error) {
	err := db.Delete(&book).Error
	if err != nil {
		return err
	}

	return nil
}