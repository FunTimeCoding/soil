package subscription

import "time"

type Subscription struct {
	Identifier        uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	Callsign          string    `gorm:"index;column:callsign"`
	RootIdentifier    string    `gorm:"index;column:root_identifier"`
	ChannelIdentifier string    `gorm:"column:channel_identifier"`
	Alias             string    `gorm:"column:alias"`
	LastEvent         time.Time `gorm:"column:last_event_at"`
	CreatedAt         time.Time `gorm:"column:created_at"`
}
