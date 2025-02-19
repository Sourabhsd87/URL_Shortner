package models

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	LongURL  string `gorm:"not null"`
	ShortURL string `gorm:"uniqueIndex"`
	Clicks   int    `gorm:"default:0"`
	Expiry   time.Time
}
