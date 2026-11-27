package mock_client

import "github.com/funtimecoding/soil/pkg/browser/ladybird/types/response"

func (c *Client) AddContext(
	identifier string,
	locator string,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.contexts = append(c.contexts, response.NewContext(identifier, locator))
}
