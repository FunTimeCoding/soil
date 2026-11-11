package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/send_result"
	"strings"
)

func sent(
	to []string,
	r *send_result.Result,
) string {
	if r.Requested != r.Delivery {
		return "Broadcast sent, queued - a broadcast never wakes anyone; name the recipients to wake them"
	}

	result := "Broadcast sent"

	if len(to) > 0 {
		result = fmt.Sprintf("Sent to %s", strings.Join(to, ", "))
	}

	if len(r.Held) == 0 {
		return result
	}

	return fmt.Sprintf(
		"%s; deferred to the next prompt of %s - immediate delivery is rate limited for them right now",
		result,
		strings.Join(r.Held, ", "),
	)
}
