package mock_client

import "github.com/funtimecoding/soil/pkg/band/response"

func (c *Client) Setup() (*response.Setup, error) {
	return response.NewSetup(), nil
}
