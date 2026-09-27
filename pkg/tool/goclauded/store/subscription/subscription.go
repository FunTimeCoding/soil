package subscription

import "time"

type Subscription struct {
	Name            string    `gorm:"primaryKey;column:name"`
	Kinds           string    `gorm:"column:kinds"`
	EventIdentifier uint      `gorm:"column:event_identifier"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	LastSeen        time.Time `gorm:"column:last_seen"`
}
