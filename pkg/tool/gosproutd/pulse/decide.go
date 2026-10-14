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

	return constant.DeliverySend
}
