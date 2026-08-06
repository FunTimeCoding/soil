package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListAddressRanges() string {
	result, e := c.client.ListAddressRanges(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
