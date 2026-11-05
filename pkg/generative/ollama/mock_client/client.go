package mock_client

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama/chat_request"
	"github.com/ollama/ollama/api"
)

type Client struct {
	Text         string
	Respond      func(prompt string) string
	Prompts      []string
	Chats        []*chat_request.Request
	Vector       []float32
	DoubleVector []float64
	Models       []api.ListModelResponse
	Processes    []api.ProcessModelResponse
	Shown        *api.ShowResponse
	Release      string
	Address      string
	Failure      error
	Closed       int
}
