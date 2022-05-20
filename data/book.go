package data

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id       uint   `gorm:"primaryKey"`
	Name     string `gorm:"UNIQUE_INDEX;type:varchar(200);not null"`
	Quantity int
}
