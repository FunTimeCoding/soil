package cruise_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/cruise"
	"time"
)

func Settled(now time.Time) *cruise.Reading {
	r := cruise.New()
	r.StateKnown = true
	r.Idle = true
	r.Grace = 30 * time.Minute
	r.IdleSince = now.Add(-31 * time.Minute)

	return r
}
