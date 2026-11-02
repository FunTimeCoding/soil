package mock_client

import "github.com/funtimecoding/soil/pkg/band/response"

func (c *Client) GeneralSettings() (*response.GeneralSettings, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return nil, c.failure
	}

	return response.NewGeneralSettings(), nil
}
