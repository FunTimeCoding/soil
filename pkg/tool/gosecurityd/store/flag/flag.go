package flag

import "time"

type Flag struct {
	Identifier    uint      `gorm:"primarykey;column:id"`
	Vulnerability string    `gorm:"not null;uniqueIndex"`
	Note          string    `gorm:"type:text"`
	Severity      string    `gorm:"not null;default:''"`
	Score         float64   `gorm:"not null;default:0"`
	Mitigated     bool      `gorm:"not null;default:false;index"`
	Mitigation    string    `gorm:"type:text"`
	CreatedAt     time.Time `gorm:"not null"`
	UpdatedAt     time.Time `gorm:"not null"`
}
