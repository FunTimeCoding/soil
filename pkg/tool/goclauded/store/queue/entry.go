package queue

import "time"

type Entry struct {
	Identifier        uint       `gorm:"primaryKey;autoIncrement;column:identifier"`
	SessionIdentifier string     `gorm:"column:session_identifier"`
	Callsign          string     `gorm:"column:callsign"`
	Kind              string     `gorm:"column:kind"`
	Body              string     `gorm:"column:body"`
	Consumed          bool       `gorm:"column:consumed"`
	ConsumedAt        *time.Time `gorm:"column:consumed_at"`
	CreatedAt         time.Time  `gorm:"column:created_at"`
}
