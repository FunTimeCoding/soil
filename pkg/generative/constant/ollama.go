package constant

import "github.com/funtimecoding/soil/pkg/web/constant"

const (
	OllamaHostEnvironment string = "OLLAMA_HOST"
	OllamaPortEnvironment string = "OLLAMA_PORT"

	OllamaHost     = constant.Localhost
	OllamaPort int = 11434

	Llama31   = "llama3.1"    // 8b
	Llama32   = "llama3.2"    // 3b
	Llama321b = "llama3.2:1b" // 1b

	OllamaEmbedModel = "nomic-embed-text" // 768 dimensions

	OllamaSystemRole    = "system"
	OllamaUserRole      = "user"
	OllamaAssistantRole = "assistant"

	OllamaNotationFormat = "json"

	OllamaContextSize = "num_ctx"
	OllamaPredictSize = "num_predict"
	OllamaTemperature = "temperature"
)
