package mock_client

import "github.com/funtimecoding/soil/pkg/band/response"

func (c *Client) ScreenService() (*response.ScreenService, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	result := response.NewScreenService()
	result.EnabledState = c.screenService

	return result, nil
}
