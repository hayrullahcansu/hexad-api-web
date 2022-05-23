package repo

import (
	"errors"
	"fmt"
	"lib-api/data"

	"gorm.io/gorm"
)

type ILibraryRepository interface {
	GetBooks() []data.Book
	BorrowOrReturnBook(action, user, book string) (data.Borrow, error)
	MyBorrowedList(user string) []data.Borrow
}

type LibraryRepository struct {
	ILibraryRepository
	db *gorm.DB
}

func NewLibraryRepository(db *gorm.DB) ILibraryRepository {
	return &LibraryRepository{
		db: db,
	}
}

func (br *LibraryRepository) GetBooks() []data.Book {
	var books []data.Book
	db := Instance()
	db.Where("quantity > ?", 0).Find(&books)
	return books
}

func (br *LibraryRepository) MyBorrowedList(user string) []data.Borrow {
	var borrows []data.Borrow
	db := Instance()
	db.Where("user = ?", user).Find(&borrows)
	return borrows
}

func (br *LibraryRepository) BorrowOrReturnBook(action, user, book string) (data.Borrow, error) {
	if action == "borrow" {
		return br.borrowBook(user, book)
	} else {
		return br.returnBook(user, book)
	}
}

func (br *LibraryRepository) borrowBook(user, book string) (data.Borrow, error) {
	db := Instance()
	var b data.Book
	var borrow data.Borrow

	err := db.Where("name = ? AND quantity > ?", book, 0).First(&b).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return borrow, errors.New(fmt.Sprintf("you cannot borrow %s book", book))
	}

	query := db.Where("user = ? AND name = ?", user, book).First(&borrow)
	if query.RowsAffected > 0 {
		return borrow, errors.New(fmt.Sprintf("you cannot borrow %s book", book))
	}
	var quantity int64
	db.Where("user = ?", user).Count(&quantity)
	if quantity > 1 {
		return borrow, errors.New(fmt.Sprintf("you cannot borrow %s book", book))
	}

	borrow.User = user
	borrow.Name = b.Name
	b.Quantity--
	db.Create(borrow)
	db.Model(&b).Where("name = ?", book).Update("quantity", b.Quantity)
	return borrow, nil
}

func (br *LibraryRepository) returnBook(user, book string) (data.Borrow, error) {
	db := Instance()
	var b data.Book
	var borrow data.Borrow

	query := db.Where("user = ? AND name = ?", user, book).First(&borrow)
	if query.RowsAffected == 0 {
		return borrow, errors.New(fmt.Sprintf("you cannot return %s book, because you don't have", book))
	}

	err := db.Where("name = ?", book).First(&b).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return borrow, errors.New(fmt.Sprintf("you cannot return %s book, because the book is not registered this library ", book))
	}

	b.Quantity++
	db.Where("user = ? AND name = ?", user, book).Delete(&borrow)
	db.Model(&b).Where("name = ?", book).Update("quantity", b.Quantity)
	return borrow, nil
}
