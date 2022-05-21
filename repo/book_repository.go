package repo

import (
	"lib-api/data"

	"gorm.io/gorm"
)

type IBookRepository interface {
	GetBooks() []data.Book
}

type BookRepository struct {
	IBookRepository
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &BookRepository{
		db: db,
	}
}

func (br *BookRepository) GetBooks() []data.Book {
	var books []data.Book
	db := Instance()
	db.Where("quantity > ?", 0).Find(&books)
	return books
}
