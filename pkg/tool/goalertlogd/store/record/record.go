package record

import "time"

type Record struct {
	Identifier  uint              `gorm:"primaryKey;autoIncrement;column:identifier"`
	Fingerprint string            `gorm:"index;column:fingerprint"`
	Name        string            `gorm:"index;column:name"`
	Severity    string            `gorm:"column:severity"`
	Summary     string            `gorm:"column:summary"`
	Labels      map[string]string `gorm:"serializer:json;column:labels"`
	Start       time.Time         `gorm:"index;column:started_at"`
	End         *time.Time        `gorm:"column:ended_at"`
}
