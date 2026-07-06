package ollama

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama/constant"
	"github.com/funtimecoding/soil/pkg/generative/ollama/generate_request"
	"github.com/funtimecoding/soil/pkg/generative/ollama/generate_response"
	"github.com/ollama/ollama/api"
)

func (c *Client) Generate(v *generate_request.Request) (*generate_response.Response, error) {
	r := v.Get()

	if r.Model == "" {
		r.Model = constant.Llama31
	}

	r.Stream = new(false)
	var result *api.GenerateResponse
	e := c.client.Generate(
		c.context,
		r,
		func(r api.GenerateResponse) error {
			result = &r

			return nil
		},
	)

	if e != nil {
		return nil, e
	}

	return generate_response.New(result), nil
}
