package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetSessionCruise(
	x context.Context,
	identifier string,
) (*client.CruiseResponse, error) {
	r, e := c.generated.GetSessionCruiseWithResponse(x, identifier)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get session cruise", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
