package models

// import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;not null;unique"`
	Password string `gorm:"size:255;not null"`
	IsAdmin  int    `gorm:"default:0"`
}
