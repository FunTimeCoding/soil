package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetSessionToolContext(
	x context.Context,
	identifier string,
	params *client.GetSessionToolContextParams,
) (*client.ToolContextResponse, error) {
	r, e := c.generated.GetSessionToolContextWithResponse(x, identifier, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"get session tool context",
			r.HTTPResponse,
			r.Body,
		)
	}

	return r.JSON200, nil
}
