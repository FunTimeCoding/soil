package digest

import (
	"github.com/funtimecoding/soil/pkg/watch/digest/event"
	"strings"
)

func counts(events []*event.Event) string {
	var order []string
	count := map[string]int{}

	for _, e := range events {
		if count[e.Noun] == 0 {
			order = append(order, e.Noun)
		}

		count[e.Noun]++
	}

	part := make([]string, 0, len(order))

	for _, noun := range order {
		part = append(part, plural(count[noun], noun))
	}

	return strings.Join(part, ", ")
}
