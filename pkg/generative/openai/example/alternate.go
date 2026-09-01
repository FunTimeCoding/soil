package example

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/sashabaranov/go-openai"
)

func Alternate() {
	c := openai.NewClient(environment.Required(constant.OpenAITokenEnvironment))
	r, e := c.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleUser, Content: "Hello!"},
			},
		},
	)
	errors.PanicOnError(e)
	console.Line(r.Choices[0].Message.Content)
}
