package entities

import "time"

type Product struct {
	ID         uint            `gorm:"primaryKey;autoIncrement"`
	Code       string          `gorm:"type:char(7);unique;not null"`
	Name       string          `gorm:"type:varchar(100);not null;uniqueIndex:uni_products_name_category,priority:1"`
	CategoryID uint            `gorm:"not null;uniqueIndex:uni_products_name_category,priority:2"`
	Category   ProductCategory `gorm:"foreignKey:CategoryID;references:ID"`
	Price      float64         `gorm:"type:numeric(10,2)"`
	CreatedAt  time.Time       `gorm:"default:now()"`
	UpdatedAt  *time.Time
	DeletedAt  *time.Time `gorm:"index"`
}
