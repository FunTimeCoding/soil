package queue

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

type Entry struct {
	Identifier        uint              `gorm:"primaryKey;autoIncrement;column:identifier"`
	SessionIdentifier string            `gorm:"column:session_identifier"`
	Callsign          string            `gorm:"column:callsign"`
	Kind              string            `gorm:"column:kind"`
	Body              string            `gorm:"column:body"`
	MessageIdentifier *uint             `gorm:"column:message_identifier"`
	Delivery          constant.Delivery `gorm:"column:delivery"`
	Immediate         bool              `gorm:"column:immediate"`
	ImmediateAt       *time.Time        `gorm:"column:immediate_at"`
	Consumed          bool              `gorm:"column:consumed"`
	ConsumedAt        *time.Time        `gorm:"column:consumed_at"`
	CreatedAt         time.Time         `gorm:"column:created_at"`
}
