package models

import "time"

type Skuf struct {
	ID        int64      `gorm:"primaryKey"`
	UserID    int64      `gorm:"uniqueIndex"`
	Name      string     `gorm:"default:Skuf"`
	Weight    float64    `gorm:"not null"`
	Alias     string     `gorm:"type:varchar(255)"`
	LastFedAt *time.Time `gorm:"column:last_fed_at"`
	FeedCount int        `gorm:"column:feed_count"`
}
