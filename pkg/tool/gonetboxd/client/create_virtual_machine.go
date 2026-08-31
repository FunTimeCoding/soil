package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateVirtualMachine(
	name string,
	cluster string,
) *response.Response {
	result, e := c.client.CreateVirtualMachine(
		c.context,
		client.CreateVirtualMachineRequest{Name: name, Cluster: cluster},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
