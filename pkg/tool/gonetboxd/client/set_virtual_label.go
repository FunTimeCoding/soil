package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) SetVirtualLabel(
	machine string,
	key string,
	value string,
) string {
	result, e := c.client.SetVirtualLabel(
		c.context,
		machine,
		key,
		client.SetVirtualLabelJSONRequestBody{Value: value},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
