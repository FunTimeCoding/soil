package client

import "github.com/funtimecoding/soil/pkg/web/locator"

func (c *Client) Address() string {
	return locator.NewOrigin(c.base).Address()
}
