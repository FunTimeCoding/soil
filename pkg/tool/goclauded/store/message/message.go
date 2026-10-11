package message

import "time"

type Message struct {
	Identifier uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	FromName   string    `gorm:"column:from_name"`
	ToName     string    `gorm:"column:to_name"`
	Body       string    `gorm:"column:body"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}
