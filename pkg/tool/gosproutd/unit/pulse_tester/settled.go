package pulse_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/pulse"
	"time"
)

func Settled(now time.Time) *pulse.Reading {
	r := pulse.New()
	r.Unpulsed = 1
	r.OldestAt = now.Add(-constant.QuietWindow - time.Minute)
	r.NewestAt = now.Add(-constant.QuietWindow - time.Minute)
	r.StateKnown = true
	r.Idle = true
	r.IdleSince = now.Add(-time.Minute)
	r.Mode = constant.CruiseOn

	return r
}
