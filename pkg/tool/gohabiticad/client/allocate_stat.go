package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) AllocateStat(stat string) *response.Response {
	result, e := c.client.AllocateStat(
		c.context,
		client.AllocateStatParamsStat(stat),
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
