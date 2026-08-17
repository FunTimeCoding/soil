package context_load

import "time"

type Load struct {
	Identifier        uint   `gorm:"primaryKey;autoIncrement;column:identifier"`
	SessionIdentifier string `gorm:"index;uniqueIndex:context_load_replay;column:session_identifier"`
	CallIdentifier    string `gorm:"uniqueIndex:context_load_replay;column:call_identifier"`
	Reference         string `gorm:"uniqueIndex:context_load_replay;column:reference"`
	Kind              string `gorm:"column:kind"`
	Name              string `gorm:"column:name"`
	Tier              string `gorm:"column:tier"`
	Query             string `gorm:"column:query"`
	OccurredAt        time.Time `gorm:"column:occurred_at"`
}
