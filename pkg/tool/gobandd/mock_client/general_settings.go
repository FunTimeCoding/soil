package mock_client

import "github.com/funtimecoding/soil/pkg/band/response"

func (c *Client) GeneralSettings() (*response.GeneralSettings, error) {
	return response.NewGeneralSettings(), nil
}
