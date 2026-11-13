package conversation_page

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/message"
	"maragu.dev/gomponents"
)

func Messages(messages []message.Message) gomponents.Node {
	var result []gomponents.Node

	for i := range messages {
		if messages[i].IsMeta {
			continue
		}

		result = append(result, messageBlock(&messages[i]))
	}

	return gomponents.Group(result)
}
