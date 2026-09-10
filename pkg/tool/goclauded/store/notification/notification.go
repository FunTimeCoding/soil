package notification

import "time"

type Notification struct {
	Identifier        uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	SessionIdentifier string    `gorm:"column:session_identifier"`
	Callsign          string    `gorm:"column:callsign"`
	Source            string    `gorm:"column:source"`
	Body              string    `gorm:"column:body"`
	Consumed          bool      `gorm:"column:consumed"`
	CreatedAt         time.Time `gorm:"column:created_at"`
}
