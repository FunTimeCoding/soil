package record

import "time"

type Record struct {
	Identifier  uint       `gorm:"primaryKey;autoIncrement;column:identifier"`
	Serial      string     `gorm:"uniqueIndex;column:serial"`
	Kind        string     `gorm:"index;column:kind"`
	Name        string     `gorm:"index;column:name"`
	CommonName  string     `gorm:"index;column:common_name"`
	Issuer      string     `gorm:"index;column:issuer"`
	Certificate string     `gorm:"column:certificate"`
	Key         string     `gorm:"column:private_key"`
	Start       time.Time  `gorm:"column:not_before"`
	End         time.Time  `gorm:"index;column:not_after"`
	Revoked     *time.Time `gorm:"column:revoked_at"`
	Published   *time.Time `gorm:"column:published_at"`
}
