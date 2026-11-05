package mock_client

import "github.com/ollama/ollama/api"

func (c *Client) List() ([]api.ListModelResponse, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.Models, nil
}
