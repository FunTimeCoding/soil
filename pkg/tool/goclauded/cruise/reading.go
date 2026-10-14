package cruise

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

type Reading struct {
	Closed     bool
	StateKnown bool
	Idle       bool
	IdleSince  time.Time
	Mode       constant.Cruise
	Pace       time.Duration
	PromotedAt time.Time
}
