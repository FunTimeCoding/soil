package basic

import "github.com/funtimecoding/soil/pkg/web/locator"

func (c *Client) Base() *locator.Locator {
	return c.base
}
