package completion

import "time"

type Completion struct {
	Identifier        uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	SessionIdentifier string    `gorm:"column:session_identifier"`
	Name              string    `gorm:"column:name"`
	Kind              string    `gorm:"column:kind"`
	Topic             string    `gorm:"column:topic"`
	Summary           string    `gorm:"column:summary"`
	Sequence          int       `gorm:"column:sequence"`
	CreatedAt         time.Time `gorm:"column:created_at"`
	CompletedAt       time.Time `gorm:"column:completed_at"`
}
