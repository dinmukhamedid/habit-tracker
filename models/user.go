package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	Age      int
	Role     string  `gorm:"default:user"`
	Habits   []Habit `gorm:"foreignKey:UserID"` // One-to-Many
}
