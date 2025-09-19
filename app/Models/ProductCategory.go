package models

type ProductCategory struct {
	ID         uint `gorm:"primaryKey"`
	ProductID  uint `gorm:"not null"`
	CategoryID uint `gorm:"not null"`
}
