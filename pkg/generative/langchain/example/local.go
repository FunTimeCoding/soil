package example

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/ollama/constant"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func Local() {
	c, newFail := ollama.New(ollama.WithModel(constant.Llama31))
	errors.PanicOnError(newFail)
	query := "very briefly, tell me the difference between a comet and a meteor"
	x := context.Background()

	if false {
		_, streamFail := llms.GenerateFromSinglePrompt(
			x,
			c,
			query,
			llms.WithStreamingFunc(
				func(
					_ context.Context,
					chunk []byte,
				) error {
					fmt.Printf("chunk len=%d: %s\n", len(chunk), chunk)

					return nil
				},
			),
		)
		errors.PanicOnError(streamFail)
	}

	response, generateFail := llms.GenerateFromSinglePrompt(
		x,
		c,
		query,
		llms.WithTemperature(0.0), // less is more deterministic
	)
	errors.PanicOnError(generateFail)
	fmt.Printf("Response: %s\n", response)
}
