package store

import "time"

type Click struct {
	Identifier uint      `gorm:"primaryKey;column:id"`
	Label      string    `gorm:"not null;index"`
	CreatedAt  time.Time `gorm:"index"`
}
