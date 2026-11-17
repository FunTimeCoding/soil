package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) SetInterfaceEthernetPower(
	x context.Context,
	device string,
	name string,
	mode string,
	kind string,
) (client.Interface, error) {
	body := client.EthernetPowerRequest{Mode: mode}

	if kind != "" {
		body.Type = &kind
	}

	r, e := c.client.SetInterfaceEthernetPowerWithResponse(
		x,
		device,
		name,
		body,
	)

	if e != nil {
		return client.Interface{}, e
	}

	if r.JSON200 == nil {
		return client.Interface{}, refusal.New(
			"set interface ethernet power",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
