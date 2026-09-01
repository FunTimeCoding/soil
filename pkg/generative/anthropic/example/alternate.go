package example

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/liushuangls/go-anthropic/v2"
)

func Alternate() {
	// https://github.com/liushuangls/go-anthropic
	c := anthropic.NewClient(
		environment.Required(constant.AnthropicTokenEnvironment),
	)
	r, e := c.CreateMessages(
		context.Background(),
		anthropic.MessagesRequest{
			Model: anthropic.ModelClaudeHaiku4Dot5,
			Messages: []anthropic.Message{
				anthropic.NewUserTextMessage("What is your name?"),
			},
			MaxTokens: 1000,
		},
	)

	if e != nil {
		if n, okay := errors.AsType[*anthropic.APIError](e); okay {
			console.Format(
				"Messages error, type: %s, message: %s",
				n.Type,
				n.Message,
			)
		} else {
			console.Format("Messages error: %v\n", e)
		}

		return
	}

	console.Line(r.Content[0].Text)
}
