package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) SetHost(
	identifier string,
	body client.HostRequest,
	apply *bool,
) (string, int) {
	result, e := c.client.SetHost(
		c.context,
		identifier,
		&client.SetHostParams{Apply: apply},
		body,
	)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
