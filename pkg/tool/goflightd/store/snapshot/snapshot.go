package snapshot

import "time"

type Snapshot struct {
	Identifier uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	Time       time.Time `gorm:"index;column:time"`
	Kind       string    `gorm:"index;column:kind"`
	Key        string    `gorm:"index;column:key"`
	Value      string    `gorm:"column:value"`
}
