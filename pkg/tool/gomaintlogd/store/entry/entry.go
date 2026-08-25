package entry

import "time"

type Entry struct {
	Identifier  uint      `gorm:"primarykey;column:id" json:"id"`
	Timestamp   time.Time `json:"timestamp" gorm:"not null;index"`
	Action      string    `json:"action" gorm:"not null;type:text"`
	User        string    `json:"user" gorm:"not null;index"`
	System      string    `json:"system" gorm:"index"`
	Service     string    `json:"service" gorm:"index"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
