package cruise

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"time"
)

type Cruise struct {
	Session    string          `gorm:"primaryKey;column:session"`
	Mode       constant.Cruise `gorm:"column:mode;not null"`
	Pace       time.Duration   `gorm:"column:pace"`
	PromotedAt *time.Time      `gorm:"column:promoted_at"`
	UpdatedAt  time.Time       `gorm:"column:updated_at"`
}
