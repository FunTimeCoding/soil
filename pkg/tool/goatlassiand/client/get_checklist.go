package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) GetChecklist(key string) string {
	result, e := c.client.GetChecklist(c.context, key)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
