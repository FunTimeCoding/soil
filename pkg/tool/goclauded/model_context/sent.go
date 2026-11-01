package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"strings"
)

func sent(
	to []string,
	held []string,
	d constant.Delivery,
) string {
	if len(to) == 0 && d != constant.DeliveryQueued {
		return "Broadcast sent, queued - a broadcast never wakes anyone; name the recipients to wake them"
	}

	result := "Broadcast sent"

	if len(to) > 0 {
		result = fmt.Sprintf("Sent to %s", strings.Join(to, ", "))
	}

	if len(held) == 0 {
		return result
	}

	if d == constant.DeliveryCruise {
		return fmt.Sprintf(
			"%s; held for the cruise setting of %s - it wakes them when the setting allows, or arrives with their next prompt",
			result,
			strings.Join(held, ", "),
		)
	}

	return fmt.Sprintf(
		"%s; deferred to the next prompt of %s - immediate delivery is rate limited for them right now",
		result,
		strings.Join(held, ", "),
	)
}
