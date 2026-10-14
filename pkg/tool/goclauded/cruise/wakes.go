package cruise

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

func Wakes(
	r *Reading,
	now time.Time,
) bool {
	if r.Closed || r.Mode == constant.CruiseOff {
		return false
	}

	if !r.StateKnown || !r.Idle {
		return false
	}

	if now.Sub(r.IdleSince) >= constant.WakeWindow {
		return false
	}

	if r.Mode == constant.CruisePaced && !r.PromotedAt.IsZero() && now.Sub(
		r.PromotedAt,
	) < r.Pace {
		return false
	}

	return true
}
