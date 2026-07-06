package example

import (
	"context"
	"fmt"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func Official() {
	c := anthropic.NewClient(
		option.WithAPIKey(environment.Required(constant.TokenEnvironment)),
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
	fmt.Printf("%+v\n", r.Content)
}
