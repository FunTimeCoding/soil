package cruise_setting

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

type Setting struct {
	Identifier        uint            `gorm:"primaryKey;autoIncrement;column:identifier"`
	SessionIdentifier string          `gorm:"uniqueIndex;column:session_identifier"`
	Mode              constant.Cruise `gorm:"column:mode;not null"`
	Pace              time.Duration   `gorm:"column:pace"`
	PromotedAt        *time.Time      `gorm:"column:promoted_at"`
	UpdatedAt         time.Time       `gorm:"column:updated_at"`
}
