package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetSessionLabel(
	x context.Context,
	identifier string,
) (*client.SessionLabelResponse, error) {
	r, e := c.generated.GetSessionLabelWithResponse(x, identifier)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get session label", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
