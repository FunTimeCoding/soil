package example

import (
	"context"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func Official() {
	c := anthropic.NewClient(
		option.WithAPIKey(
			environment.Required(constant.AnthropicTokenEnvironment),
		),
	)
	r, e := c.Messages.New(
		context.Background(),
		anthropic.MessageNewParams{
			MaxTokens: 1024,
			Messages: []anthropic.MessageParam{
				{
					Role: anthropic.MessageParamRoleUser,
					Content: []anthropic.ContentBlockParamUnion{
						{
							OfText: &anthropic.TextBlockParam{
								Text: "Explain a cat in 10 words.",
							},
						},
					},
				},
			},
			Model: anthropic.ModelClaudeSonnet4_5,
		},
	)
	errors.PanicOnError(e)
	console.Format("%+v\n", r.Content)
}
