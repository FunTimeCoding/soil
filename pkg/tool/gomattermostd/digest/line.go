package digest

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
)

func line(e *event.Event) string {
	switch e.Kind {
	case constant.ReactionAddedEvent:
		return fmt.Sprintf("%s reacted %s", e.Author, e.Text)
	case constant.ReactionRemovedEvent:
		return fmt.Sprintf("%s removed %s", e.Author, e.Text)
	default:
		return fmt.Sprintf("%s: %s", e.Author, excerpt(e.Text))
	}
}
