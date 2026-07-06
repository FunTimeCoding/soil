package ollama

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama/constant"
	"github.com/ollama/ollama/api"
)

func (c *Client) GenerateStream(
	r *api.GenerateRequest,
	output chan<- string,
	response *api.GenerateResponse,
) error {
	if r.Model == "" {
		r.Model = constant.Llama31
	}

	r.Stream = new(true)
	e := c.client.Generate(
		c.context,
		r,
		func(r api.GenerateResponse) error {
			if r.Done {
				response.Response = r.Response
				response.Metrics = r.Metrics
				response.Done = r.Done
				response.Model = r.Model
				response.CreatedAt = r.CreatedAt
			}

			output <- r.Response

			return nil
		},
	)
	close(output)

	return e
}
