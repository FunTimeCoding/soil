package cruise_setting

import "time"

type Setting struct {
	Identifier        uint           `gorm:"primaryKey;autoIncrement;column:identifier"`
	SessionIdentifier string         `gorm:"uniqueIndex;column:session_identifier"`
	Grace             *time.Duration `gorm:"column:grace"`
	Woken             int            `gorm:"column:woken;not null;default:0"`
	WokenAt           *time.Time     `gorm:"column:woken_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at"`
}
