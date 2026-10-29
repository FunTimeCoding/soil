package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateTenant(
	x context.Context,
	name string,
) (client.Tenant, error) {
	r, e := c.client.CreateTenantWithResponse(
		x,
		client.CreateNameRequest{Name: name},
	)

	if e != nil {
		return client.Tenant{}, e
	}

	if r.JSON201 == nil {
		return client.Tenant{}, refusal.New(
			"create tenant",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
