package mock_client

import "github.com/funtimecoding/soil/pkg/band/response"

func (c *Client) Redirection() (*response.Redirection, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	result := response.NewRedirection()
	result.ListenerEnabled = c.listener

	return result, nil
}
