package models

import "time"

type URL struct {
	ID          uint      `gorm:"primaryKey"`
	ShortCode   string    `gorm:"uniqueIndex;size:20"`
	OriginalURL string    `gorm:"type:text;not null"`
	CreatedAt   time.Time
}
