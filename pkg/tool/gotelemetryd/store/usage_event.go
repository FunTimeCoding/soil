package store

import "time"

type UsageEvent struct {
	Identifier          uint   `gorm:"primaryKey;column:id"`
	Tool                string `gorm:"not null;index"`
	Surface             string `gorm:"not null;index"`
	Actor               string `gorm:"not null;index"`
	Outcome             string `gorm:"not null"`
	Kind                string `gorm:"not null;index"`
	DurationMillisecond int64
	Detail              *string   `gorm:"type:jsonb"`
	CreatedAt           time.Time `gorm:"index"`
}
