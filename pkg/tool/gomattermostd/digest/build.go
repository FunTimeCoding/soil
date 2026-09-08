package digest

import (
	"fmt"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
)

func Build(alias string, events []*event.Event) string {
	if len(events) == 0 {
		return ""
	}

	if len(events) == 1 {
		return fmt.Sprintf(
			"[%s] %s · %s",
			alias,
			events[0].At.Format(timeConstant.HourMinute),
			line(events[0]),
		)
	}

	if len(events) <= constant.EnumerateLimit {
		result := enumerated(alias, events)

		if len([]rune(result)) <= constant.DigestBudget {
			return result
		}
	}

	return collapsed(alias, events)
}
