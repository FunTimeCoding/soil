package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) RemoveVirtualLabel(
	machine string,
	key string,
) {
	result, e := c.client.RemoveVirtualLabel(c.context, machine, key)
	errors.PanicOnError(e)
	web.ReadString(result)
}
