package mark

import "time"

type Mark struct {
	Identifier uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	Time       time.Time `gorm:"index;column:time"`
	Label      string    `gorm:"column:label"`
	Note       string    `gorm:"column:note"`
}
