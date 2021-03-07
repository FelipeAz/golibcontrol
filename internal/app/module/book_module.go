package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/FelipeAz/golibcontrol/internal/app/repository"
)

// BookModule process the request recieved from handler.
type BookModule struct {
	repository repository.BookRepository
}

// GetBooks returns all books on DB.
func (m BookModule) GetBooks() (books []model.Book, err error) {
	books, err = m.repository.GetBooks()
	return
}

// GetBook returns all books on DB.
func (m BookModule) GetBook(id int) (book model.Book, err error) {
	book, err = m.repository.GetBook(id)
	return
}

// CreateBook persist a book to the database.
func (m BookModule) CreateBook(book model.Book) (id uint, err error) {
	id, err = m.repository.CreateBook(book)
	return
}

// UpdateBook update an existent book.
func (m BookModule) UpdateBook(id int, upBook model.Book) (book model.Book, err error) {
	book, err = m.repository.UpdateBook(id, upBook)
	return
}

// DeleteBook delete an existent book.
func (m BookModule) DeleteBook(id int) (bool, error) {
	err := m.repository.DeleteBook(id)
	if err != nil {
		return false, err
	}

	return true, nil
}