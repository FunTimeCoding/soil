package frame

import "time"

type Frame struct {
	Identifier uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	Session    string    `gorm:"column:session;not null;uniqueIndex:frame_session_name"`
	Name       string    `gorm:"column:name;not null;uniqueIndex:frame_session_name"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}
