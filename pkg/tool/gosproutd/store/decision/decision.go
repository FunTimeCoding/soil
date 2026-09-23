package decision

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/choice"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/frame"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/turn"
	"time"
)

type Decision struct {
	Identifier    uint                   `gorm:"primaryKey;autoIncrement;column:identifier"`
	Session       string                 `gorm:"column:session;not null;index"`
	Question      string                 `gorm:"column:question;not null"`
	DefaultAction string                 `gorm:"column:default_action;not null"`
	State         constant.State         `gorm:"column:state;not null;index"`
	AnswerKind    constant.AnswerKind    `gorm:"column:answer_kind"`
	AnswerChannel constant.AnswerChannel `gorm:"column:answer_channel"`
	Answer        string                 `gorm:"column:answer"`
	Resolution    constant.Resolution    `gorm:"column:resolution"`
	ClearLine     string                 `gorm:"column:clear_line"`
	Bumped        int                    `gorm:"column:bumped;not null"`
	Frames        []*frame.Frame         `gorm:"many2many:decision_frame;"`
	Choices       []*choice.Choice       `gorm:"foreignKey:DecisionIdentifier"`
	Turns         []*turn.Turn           `gorm:"foreignKey:DecisionIdentifier"`
	CreatedAt     time.Time              `gorm:"column:created_at"`
	UpdatedAt     time.Time              `gorm:"column:updated_at"`
	AnsweredAt    *time.Time             `gorm:"column:answered_at"`
	SeenAt        *time.Time             `gorm:"column:seen_at"`
	PulsedAt      *time.Time             `gorm:"column:pulsed_at"`
}
