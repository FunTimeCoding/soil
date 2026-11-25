package digest

import (
	"fmt"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/watch/constant"
	"github.com/funtimecoding/soil/pkg/watch/digest/event"
)

func Build(
	alias string,
	events []*event.Event,
) string {
	if len(events) == 0 {
		return ""
	}

	if len(events) == 1 {
		return fmt.Sprintf(
			"[%s] %s · %s",
			alias,
			events[0].At.Format(timeConstant.HourMinute),
			events[0].Line,
		)
	}

	if len(events) <= constant.EnumerateLimit {
		result := enumerated(alias, events)

		if len([]rune(result)) <= constant.DigestLimit {
			return result
		}
	}

	return collapsed(alias, events)
}
