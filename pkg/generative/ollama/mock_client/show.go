package mock_client

import "github.com/ollama/ollama/api"

func (c *Client) Show(_ string) (*api.ShowResponse, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.Shown, nil
}
