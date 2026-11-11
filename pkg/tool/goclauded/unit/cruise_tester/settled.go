package cruise_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/cruise_reading"
	"time"
)

func Settled(now time.Time) *cruise_reading.Reading {
	r := cruise_reading.New()
	r.StateKnown = true
	r.Idle = true
	r.Grace = 30 * time.Minute
	r.IdleSince = now.Add(-31 * time.Minute)

	return r
}
