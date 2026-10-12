package mock_client

import "github.com/funtimecoding/soil/pkg/band/response"

func (c *Client) Redirection() (*response.Redirection, error) {
	return response.NewRedirection(), nil
}
