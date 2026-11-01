package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostSessionCruise(
	x context.Context,
	identifier string,
	body client.PostSessionCruiseJSONRequestBody,
) (*client.CruiseResponse, error) {
	r, e := c.generated.PostSessionCruiseWithResponse(x, identifier, body)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("post session cruise", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
