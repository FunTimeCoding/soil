package store

import "time"

type Click struct {
	ID        uint      `gorm:"primaryKey"`
	Label     string    `gorm:"not null;index"`
	CreatedAt time.Time `gorm:"index"`
}
