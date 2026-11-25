package digest

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/watch/constant"
	"github.com/funtimecoding/soil/pkg/watch/digest/event"
)

func collapsed(
	alias string,
	events []*event.Event,
) string {
	return fmt.Sprintf(
		"[%s] %d new · %s · %s · %s\n%slatest — %s",
		alias,
		len(events),
		span(events),
		authors(events),
		counts(events),
		constant.DigestIndent,
		events[len(events)-1].Line,
	)
}
