package cruise_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/cruise"
	"time"
)

func Settled(now time.Time) *cruise.Reading {
	r := cruise.New()
	r.StateKnown = true
	r.Idle = true
	r.IdleSince = now.Add(-time.Minute)
	r.Mode = constant.CruiseOn

	return r
}
