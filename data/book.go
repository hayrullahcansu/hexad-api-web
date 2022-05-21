package data

type Book struct {
	Name     string `gorm:"UNIQUE_INDEX;type:varchar(200);not null"`
	Quantity int
}
