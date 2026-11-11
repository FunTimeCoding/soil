package cruise_reading

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

func (r *Reading) Wakes(now time.Time) bool {
	if r.Closed || r.Grace <= 0 {
		return false
	}

	if !r.StateKnown || !r.Idle {
		return false
	}

	idle := now.Sub(r.IdleSince)

	return idle >= r.Grace && idle < constant.WakeWindow
}
