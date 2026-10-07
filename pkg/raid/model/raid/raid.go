package raid

import "time"

type Raid struct {
	Identifier uint `gorm:"primaryKey;column:id"`
	Name       string
	Date       time.Time
	CreatedAt  time.Time
}
