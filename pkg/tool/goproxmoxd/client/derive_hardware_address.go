package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) DeriveHardwareAddress(identifier int) (string, int) {
	result, e := c.client.DeriveHardwareAddressWithResponse(
		c.context,
		&client.DeriveHardwareAddressParams{
			Instance:   &c.instance,
			Identifier: identifier,
		},
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
