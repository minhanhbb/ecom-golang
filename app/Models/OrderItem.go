package models

// import "gorm.io/gorm"

type OrderItems struct {
	ID        uint    `gorm:"primaryKey"`
	OrderId   int     `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	ProductID int     `gorm:"default:0"`
	Price     float64 `gorm:"not null"`
}
