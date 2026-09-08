package digest

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"strings"
)

func counts(events []*event.Event) string {
	message := 0
	reaction := 0

	for _, e := range events {
		if e.Kind == constant.MessageEvent {
			message++
		} else {
			reaction++
		}
	}

	var part []string

	if message > 0 {
		part = append(part, plural(message, "message"))
	}

	if reaction > 0 {
		part = append(part, plural(reaction, "reaction"))
	}

	return strings.Join(part, ", ")
}
