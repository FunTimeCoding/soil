package turn

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"time"
)

type Turn struct {
	Identifier         uint            `gorm:"primaryKey;autoIncrement;column:identifier"`
	DecisionIdentifier uint            `gorm:"column:decision_identifier;not null;index"`
	Author             constant.Author `gorm:"column:author;not null"`
	Content            string          `gorm:"column:content;not null"`
	CreatedAt          time.Time       `gorm:"column:created_at"`
}
