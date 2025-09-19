package models

// import "gorm.io/gorm"

type Product struct {
	ID         uint    `gorm:"primaryKey"`
	Name       string  `gorm:"size:255;not null"`
	Desc       string  `gorm:"size:255;not null"`
	Price      float64 `gorm:"not null"`
	CategoryID uint    `gorm:"-"` // bỏ không dùng, dùng bảng phụ
	Images     string  `gorm:"type:text"`
	Status     int     `gorm:"default:1"`
}
