package function

import (
	"encoding/json"
	"github.com/tmc/langchaingo/llms"
)

var functions = []llms.FunctionDefinition{
	{
		Name:        "getCurrentWeather",
		Description: "Get the current weather in a given location",
		Parameters: json.RawMessage(
			`{
			"type": "object",
			"properties": {
				"location": {"type": "string", "description": "The city and state, e.g. San Francisco, CA"},
				"unit": {"type": "string", "enum": ["celsius", "fahrenheit"]}
			},
			"required": ["location", "unit"]
		}`,
		),
	},
	{
		// I found that providing a tool for Ollama to give the final response significantly
		// increases the chances of success.
		Name:        "finalResponse",
		Description: "Provide the final response to the user query",
		Parameters: json.RawMessage(
			`{
			"type": "object",
			"properties": {
				"response": {"type": "string", "description": "The final response to the user query"}
			},
			"required": ["response"]
		}`,
		),
	},
}
