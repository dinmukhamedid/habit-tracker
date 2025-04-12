package models

type Habit struct {
	ID     uint `gorm:"primaryKey"`
	Title  string
	UserID uint
	User   User `gorm:"foreignKey:UserID"`
}
