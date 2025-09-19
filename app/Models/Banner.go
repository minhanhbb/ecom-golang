package models

// import "gorm.io/gorm"

type Banners  struct {
	ID       uint   `gorm:"primaryKey"`
	Desc     string `gorm:"size:255;not null"`
	Image    string `gorm:"size:255;not null"`
}
