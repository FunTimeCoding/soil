package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) AllocateStat(stat string) (string, int) {
	result, e := c.client.AllocateStat(
		c.context,
		client.AllocateStatParamsStat(stat),
	)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
