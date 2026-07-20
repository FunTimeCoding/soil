package tracker_state

import "time"

type Record struct {
	Identifier string    `gorm:"primaryKey;column:identifier"`
	Payload    string    `gorm:"column:payload"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}
