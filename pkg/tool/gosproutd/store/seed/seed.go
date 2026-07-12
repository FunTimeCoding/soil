package seed

import "time"

type Seed struct {
	Identifier  uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	Name        string    `gorm:"column:name;not null"`
	Path        string    `gorm:"column:path;not null;uniqueIndex"`
	ContentHash string    `gorm:"column:content_hash;not null"`
	Content     string    `gorm:"column:content;not null"`
	Position    int       `gorm:"column:position;not null;index"`
	ModifiedAt  time.Time `gorm:"column:modified_at"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
