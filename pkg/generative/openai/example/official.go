package example

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func Official() {
	c := openai.NewClient(
		option.WithAPIKey(environment.Required(constant.OpenAITokenEnvironment)),
	)
	r, e := c.Chat.Completions.New(
		context.Background(),
		openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage("Say this is a test"),
			},
			Model: openai.ChatModelGPT4o,
		},
	)
	errors.PanicOnError(e)
	println(r.Choices[0].Message.Content)
}
