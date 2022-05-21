package repo

import (
	"lib-api/data"

	"gorm.io/gorm"
)

func NewTestLibraryRepository() (ILibraryRepository, error) {
	db, err := mookTestDb()
	if err != nil {
		return nil, err
	}
	repo := NewLibraryRepository(db)
	return repo, nil
}

func mookTestDb() (*gorm.DB, error) {
	fileName := "file:memdb1?mode=memory&cache=shared"
	db, err := OpenDb(fileName)
	initTestData(db)
	return db, err
}

func initTestData(db *gorm.DB) {
	for _, v := range getTestBooks() {
		db.Create(&v)
	}
}

func getTestBooks() []data.Book {
	books := make([]data.Book, 0)
	books = append(books, data.Book{Name: "TestBook1", Quantity: 3})
	books = append(books, data.Book{Name: "TestBook2", Quantity: 5})
	books = append(books, data.Book{Name: "TestBook3", Quantity: 1})
	return books
}
