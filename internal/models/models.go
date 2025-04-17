package models

type Skuf struct {
	ID     int64   `gorm:"primaryKey"`
	UserID int64   `gorm:"uniqueIndex"`
	Name   string  `gorm:"default:Skuf"`
	Weight float64 `gorm:"not null"`
}
