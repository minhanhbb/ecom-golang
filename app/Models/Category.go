package models

// import "gorm.io/gorm"

type Categories  struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255;not null"`
	Image    string `gorm:"size:255;not null"`
}
