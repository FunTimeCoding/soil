package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateAddress(
	device string,
	interfaceName string,
	address string,
	status string,
) *response.Response {
	body := client.CreateAddressRequest{
		Interface: interfaceName,
		Address:   address,
	}

	if status != "" {
		body.Status = &status
	}

	result, e := c.client.CreateAddress(c.context, device, body)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
