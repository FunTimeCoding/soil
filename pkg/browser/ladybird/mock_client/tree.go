package mock_client

import "github.com/funtimecoding/soil/pkg/browser/types/response"

func (c *Client) Tree() ([]*response.Context, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return append([]*response.Context{}, c.contexts...), nil
}
