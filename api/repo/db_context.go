package repo

import (
	"fmt"
	"lib-api/data"
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var onceMigration, onceSeed sync.Once

func getDbContext() (*gorm.DB, error) {
	return OpenDb(
		"file:memdb1?mode=memory&cache=shared",
	)
}

func Instance() *gorm.DB {
	db, err := getDbContext()
	if err != nil {
		log.Fatal("cannot initialize db", err)
	}
	onceSeed.Do(func() {
		fmt.Println("database seed first")
		initData(db)
		fmt.Println("database seed done")
	})
	return db
}

func OpenDb(connection string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(connection), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	onceMigration.Do(func() {
		// Migrate the schema
		fmt.Println("migration first")
		db.AutoMigrate(
			&data.Book{},
			&data.Borrow{},
		)
		fmt.Println("migration done")

	})
	return db, err
}

func initData(db *gorm.DB) {
	for _, v := range getTestBooks() {
		db.Create(&v)
	}
}

func getBooks() []data.Book {
	books := make([]data.Book, 0)
	books = append(books, data.Book{Name: "Sapiens: A Brief History of Humankind", Quantity: 5})
	books = append(books, data.Book{Name: "Outliers: The Story of Success", Quantity: 2})
	books = append(books, data.Book{Name: "Thinking, Fast and Slow", Quantity: 1})
	return books
}
