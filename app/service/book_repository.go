package service

import "Go-Book/app/model"

type Bookservices interface {
	SetBook(book *model.Book) (id int, err error)
	GetBookList() []model.Book
	UpdateBook(book *model.Book) error
	DeleteBook(id int) error
	FindByID(id int) (book model.Book, err error)
}
