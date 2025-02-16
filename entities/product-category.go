package entities

import "time"

type ProductCategory struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(50);uniqueIndex"`
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt *time.Time
	DeletedAt *time.Time `gorm:"index"`
}
