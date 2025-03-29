package models

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"unique"`
	Email string `gorm:"unique"`
	Age   int
}
