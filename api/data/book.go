package data

type Book struct {
	Name     string `gorm:"UNIQUE_INDEX;not null"`
	Quantity int
}
