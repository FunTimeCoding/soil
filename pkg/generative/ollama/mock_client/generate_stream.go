package mock_client

import "github.com/ollama/ollama/api"

func (c *Client) GenerateStream(
	r *api.GenerateRequest,
	output chan<- string,
	response *api.GenerateResponse,
) error {
	defer close(output)
	text := c.answer(r.Prompt)

	if c.Failure != nil {
		return c.Failure
	}

	output <- text
	response.Response = text
	response.Done = true

	return nil
}
