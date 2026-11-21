package mirror

import "time"

type Mirror struct {
	Identifier uint      `gorm:"primarykey;column:id"`
	Start      int       `gorm:"not null;default:0"`
	Total      int       `gorm:"not null;default:0"`
	Filled     bool      `gorm:"not null;default:false"`
	Synced     time.Time `gorm:"not null"`
}
