package rate_snapshot

import "time"

type Snapshot struct {
	Identifier      uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	FiveHourPercent int       `gorm:"column:five_hour_percent"`
	SevenDayPercent int       `gorm:"column:seven_day_percent"`
	FiveHourReset   time.Time `gorm:"column:five_hour_reset"`
	SevenDayReset   time.Time `gorm:"column:seven_day_reset"`
	CreatedAt       time.Time `gorm:"column:created_at"`
}
