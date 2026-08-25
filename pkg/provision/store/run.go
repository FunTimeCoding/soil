package store

import "time"

type Run struct {
	Identifier          uint `gorm:"primaryKey;column:id"`
	CreatedAt           time.Time
	Scope               string
	TriggerSource       string
	DurationMillisecond int64
	Status              string
	Output              string
	ErrorOutput         string
	GitHead             string
}
