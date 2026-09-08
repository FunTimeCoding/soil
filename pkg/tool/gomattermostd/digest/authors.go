package digest

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"strings"
)

func authors(events []*event.Event) string {
	var order []string
	seen := map[string]bool{}

	for _, e := range events {
		if seen[e.Author] {
			continue
		}

		seen[e.Author] = true
		order = append(order, e.Author)
	}

	if len(order) <= 2 {
		return strings.Join(order, ", ")
	}

	return fmt.Sprintf("%s +%d", strings.Join(order[:2], ", "), len(order)-2)
}
