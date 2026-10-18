package impression_outbox

import "time"

type Impression struct {
	Identifier uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	Source     string    `gorm:"column:source"`
	Content    string    `gorm:"column:content"`
	CapturedAt time.Time `gorm:"column:captured_at"`
}
