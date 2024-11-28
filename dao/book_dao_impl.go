package dao

import (
	"gin-demo/model"
	"github.com/jinzhu/gorm"
)

type BookDaoImpl struct {
	DB *gorm.DB
}

func (b *BookDaoImpl) AddBook(book *model.Book) error {
	if createResult := b.DB.Create(book); createResult.Error != nil {
		return createResult.Error
	}
	return nil
}

func (b *BookDaoImpl) UpdateBook(book *model.Book) error {
	if saveResult := b.DB.Save(book); saveResult.Error != nil {
		return saveResult.Error
	}
	return nil
}

func (b *BookDaoImpl) DeleteBook(book *model.Book) error {
	if deleteResult := b.DB.Delete(book); deleteResult.Error != nil {
		return deleteResult.Error
	}
	return nil
}

func (b *BookDaoImpl) ListByBookId(id uint) (*model.Book, error) {
	var book model.Book
	if listResult := b.DB.Where("id = ?", id).First(&book); listResult.Error != nil {
		return nil, listResult.Error
	}
	return &book, nil
}

//func (b *BookDaoImpl) UpdateBook(book *model.Book) error {
//	result := b.DB.Model(&model.Book{}).Where("id = ?", book.ID).
//		Updates(map[string]interface{}{
//			"bookname":  book.BookName,
//			"bookprice": book.BookPrice,
//		})
//	return result.Error
//}
