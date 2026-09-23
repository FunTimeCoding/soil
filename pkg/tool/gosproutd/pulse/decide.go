package pulse

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"time"
)

func Decide(
	r *Reading,
	now time.Time,
) constant.Delivery {
	if r.Unpulsed == 0 || r.Closed {
		return constant.DeliveryDrop
	}

	if now.Sub(r.NewestAt) < constant.QuietWindow && now.Sub(
		r.OldestAt,
	) < constant.MaximumHold {
		return constant.DeliveryHold
	}

	if r.Mode == constant.CruiseOff {
		return constant.DeliveryQueue
	}

	if !r.StateKnown || !r.Idle || now.Sub(
		r.IdleSince,
	) >= constant.WakeWindow {
		return constant.DeliveryQueue
	}

	if r.Mode == constant.CruisePaced && !r.PromotedAt.IsZero() && now.Sub(
		r.PromotedAt,
	) < r.Pace {
		return constant.DeliveryHold
	}

	return constant.DeliveryImmediate
}
