package repo

import (
	"fmt"
	"lib-api/data"
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var once sync.Once

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
	return db
}

func OpenDb(connection string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(connection), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	once.Do(func() {
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
