package fable_snapshot

import "time"

type Snapshot struct {
	Identifier uint       `gorm:"primaryKey;autoIncrement;column:identifier"`
	Percent    int        `gorm:"column:percent"`
	Reset      string     `gorm:"column:reset"`
	ResetAt    *time.Time `gorm:"column:reset_at"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
}
