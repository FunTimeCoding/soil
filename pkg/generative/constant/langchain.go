package constant

import (
	"encoding/json"
	"github.com/tmc/langchaingo/llms"
)

const LangchainTemperatureKey = "temperature"

var LangchainExampleFunctions = []llms.FunctionDefinition{
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
