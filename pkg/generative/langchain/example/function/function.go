package function

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/constant"
	"github.com/spf13/pflag"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"log"
)

var flagVerbose = pflag.Bool("v", false, "verbose mode")

func Function() {
	pflag.Parse()
	c, clientFail := ollama.New(
		ollama.WithModel(constant.Llama31),
		ollama.WithFormat(constant.NotationFormat),
	)
	errors.PanicOnError(clientFail)
	var messages []llms.MessageContent
	// system message defines the available tools.
	messages = append(
		messages,
		llms.TextParts(llms.ChatMessageTypeSystem, systemMessage()),
	)
	messages = append(
		messages,
		llms.TextParts(
			llms.ChatMessageTypeHuman,
			"What's the weather like in Beijing?",
		),
	)
	x := context.Background()

	for retries := 3; retries > 0; retries -= 1 {
		resp, generateFail := c.GenerateContent(x, messages)
		errors.PanicOnError(generateFail)
		choice1 := resp.Choices[0]
		messages = append(
			messages,
			llms.TextParts(llms.ChatMessageTypeAI, choice1.Content),
		)

		if a := unmarshalCall(choice1.Content); a != nil {
			log.Printf("Call: %v", a.Tool)

			if *flagVerbose {
				log.Printf("Call: %v (raw: %v)", a.Tool, choice1.Content)
			}

			m, cont := dispatchCall(a)

			if !cont {
				break
			}

			messages = append(messages, m)
		} else {
			// Ollama doesn't always respond with a function call, let it try again.
			log.Printf("Not a call: %v", choice1.Content)
			messages = append(
				messages,
				llms.TextParts(
					llms.ChatMessageTypeHuman,
					"Sorry, I don't understand. Please try again.",
				),
			)
		}

		if retries == 1 {
			log.Fatal("retries exhausted")
		}
	}
}
