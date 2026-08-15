package event

import "time"

type Event struct {
	Identifier uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	Time       time.Time `gorm:"index;column:time"`
	Process    string    `gorm:"index;column:process"`
	Subsystem  string    `gorm:"index;column:subsystem"`
	Category   string    `gorm:"column:category"`
	Message    string    `gorm:"column:message"`
}
