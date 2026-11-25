package outgoing

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/response"
	"html"
	"strings"
)

func New(text string) *Message {
	return &Message{
		Body: response.Body{
			ContentType: constant.TeamsHypertextContent,
			Content: strings.ReplaceAll(
				html.EscapeString(text),
				"\n",
				"<br>",
			),
		},
	}
}
