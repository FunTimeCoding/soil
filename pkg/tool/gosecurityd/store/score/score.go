package score

import "time"

type Score struct {
	Identifier string    `gorm:"primaryKey;column:identifier"`
	Score      float64   `gorm:"not null;default:0"`
	Severity   string    `gorm:"not null;default:''"`
	Version    string    `gorm:"not null;default:''"`
	Modified   time.Time `gorm:"not null"`
}
