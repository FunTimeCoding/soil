package mock_client

import "github.com/ollama/ollama/api"

func (c *Client) Running() ([]api.ProcessModelResponse, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.Processes, nil
}
