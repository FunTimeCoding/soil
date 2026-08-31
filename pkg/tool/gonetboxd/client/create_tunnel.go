package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateTunnel(
	name string,
	encapsulation string,
	group string,
) *response.Response {
	result, e := c.client.CreateTunnel(
		c.context,
		client.CreateTunnelRequest{
			Name:          name,
			Encapsulation: encapsulation,
			Group:         group,
		},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
