package constant

import "github.com/funtimecoding/soil/pkg/web/constant"

const (
	HostEnvironment string = "OLLAMA_HOST"
	PortEnvironment string = "OLLAMA_PORT"

	Host     = constant.Localhost
	Port int = 11434

	Llama31   = "llama3.1"    // 8b
	Llama32   = "llama3.2"    // 3b
	Llama321b = "llama3.2:1b" // 1b

	SystemRole    = "system"
	UserRole      = "user"
	AssistantRole = "assistant"

	NotationFormat = "json"

	ContextSize = "num_ctx"
	PredictSize = "num_predict"
	Temperature = "temperature"
)
