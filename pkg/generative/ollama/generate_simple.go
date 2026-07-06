package ollama

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama/generate_request"
	"github.com/funtimecoding/soil/pkg/generative/ollama/generate_response"
)

func (c *Client) GenerateSimple(p string) *generate_response.Response {
	return c.MustGenerate(generate_request.New().Prompt(p))
}
