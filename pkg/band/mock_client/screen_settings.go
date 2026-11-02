package mock_client

import "github.com/funtimecoding/soil/pkg/band/response"

func (c *Client) ScreenSettings() (*response.ScreenSettings, error) {
	return response.NewScreenSettings(), nil
}
