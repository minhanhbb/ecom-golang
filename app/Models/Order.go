package models

// import "gorm.io/gorm"

type Orders  struct {
	ID       uint   `gorm:"primaryKey"`
	UserID     string `gorm:"size:255;not null"`
	TotalPrice    string `gorm:"size:255;not null"`
	Status     int    `gorm:"default:0"` // 0: pending, 1: completed, 2: cancelled
}
