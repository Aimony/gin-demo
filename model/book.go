package model

const BookTableName = "book"

type Book struct {
	ID        int64   `gorm:"column:id"`
	BookName  string  `gorm:"column:b_name"`
	BookPrice float64 `gorm:"column:b_price"`
}

func (b *Book) TableName() string {
	return BookTableName
}
