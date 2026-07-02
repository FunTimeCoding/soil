package function

import (
	"github.com/tmc/langchaingo/llms"
	"log"
)

func dispatchCall(c *Call) (llms.MessageContent, bool) {
	// ollama doesn't always respond with a *valid* function call. As we're using prompt
	// engineering to inject the tools, it may hallucinate.
	if !validTool(c.Tool) {
		log.Printf(
			"invalid function call: %#v, prompting model to try again",
			c,
		)

		return llms.TextParts(
			llms.ChatMessageTypeHuman,
			"Tool does not exist, please try again.",
		), true
	}

	// we could make this more dynamic by parsing the function schema
	switch c.Tool {
	case "getCurrentWeather":
		l, okay := c.Input["location"].(string)

		if !okay {
			log.Fatal("invalid input")
		}

		unit, okay := c.Input["unit"].(string)

		if !okay {
			log.Fatal("invalid input")
		}

		weather := getCurrentWeather(l, unit)

		return llms.TextParts(llms.ChatMessageTypeHuman, weather), true
	case "finalResponse":
		resp, okay := c.Input["response"].(string)

		if !okay {
			log.Fatal("invalid input")
		}

		log.Printf("Final response: %v", resp)

		return llms.MessageContent{}, false
	default:
		// we already checked above if we had a valid tool
		panic("unreachable")
	}
}
