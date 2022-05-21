package data

type Borrow struct {
	User string `gorm:"index:idx_member;not null"`
	Name string `gorm:"index:idx_member;not null"`
}
