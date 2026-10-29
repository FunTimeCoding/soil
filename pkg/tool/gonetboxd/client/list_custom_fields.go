package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListCustomFields(
	x context.Context,
) ([]*client.CustomField, error) {
	r, e := c.client.ListCustomFieldsWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list custom fields", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
