package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateVirtualAddress(
	x context.Context,
	vmName string,
	interfaceName string,
	address string,
	status string,
) (client.Address, error) {
	body := client.CreateAddressRequest{
		Interface: interfaceName,
		Address:   address,
	}

	if status != "" {
		body.Status = &status
	}

	r, e := c.client.CreateVirtualAddressWithResponse(x, vmName, body)

	if e != nil {
		return client.Address{}, e
	}

	if r.JSON201 == nil {
		return client.Address{}, refusal.New(
			"create virtual address",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
