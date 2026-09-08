package digest

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
)

func span(events []*event.Event) string {
	first := events[0].At.Format(constant.HourMinute)
	last := events[len(events)-1].At.Format(constant.HourMinute)

	if first == last {
		return first
	}

	return join.Empty(first, "-", last)
}
