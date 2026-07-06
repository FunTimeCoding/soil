package conversations

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/message"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func messageBlock(m *message.Message) gomponents.Node {
	class := "message message-user"

	if m.Role == "assistant" {
		class = "message message-assistant"
	}

	return html.Div(
		html.Class(class),
		html.Div(
			html.Class("message-role"),
			gomponents.Text(m.Role),
		),
		html.Div(
			html.Class("message-text"),
			gomponents.Text(m.Text),
		),
	)
}
