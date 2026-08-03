package function

import (
	"context"
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"log"
)

func Function() {
	a := argument.NewSimple("function")
	a.BooleanShort(argumentConstant.Verbose, "v", false, "verbose mode")
	a.ParseSimple()
	verbose := a.GetBoolean(argumentConstant.Verbose)
	c, clientFail := ollama.New(
		ollama.WithModel(constant.Llama31),
		ollama.WithFormat(constant.OllamaNotationFormat),
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

			if verbose {
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
