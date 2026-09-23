package pulse

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"time"
)

type Reading struct {
	Unpulsed   int
	OldestAt   time.Time
	NewestAt   time.Time
	Closed     bool
	Idle       bool
	IdleSince  time.Time
	StateKnown bool
	Mode       constant.Cruise
	Pace       time.Duration
	PromotedAt time.Time
}
