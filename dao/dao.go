package dao

import "gin-demo/model"

type BookDao interface {
	AddBook(book *model.Book) error
	UpdateBook(book *model.Book) error
	DeleteBook(book *model.Book) error
	ListByBookId(id uint) (*model.Book, error)
}
