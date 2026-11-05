package mock_client

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama/generate_request"
	"github.com/funtimecoding/soil/pkg/generative/ollama/generate_response"
	"github.com/ollama/ollama/api"
)

func (c *Client) Generate(
	v *generate_request.Request,
) (*generate_response.Response, error) {
	text := c.answer(v.Get().Prompt)

	if c.Failure != nil {
		return nil, c.Failure
	}

	return generate_response.New(&api.GenerateResponse{Response: text}), nil
}
