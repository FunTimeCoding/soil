package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) DeleteSessionByIdentifier(
	x context.Context,
	identifier string,
	params *client.DeleteSessionByIdentifierParams,
) (*client.DeleteReceiptResponse, error) {
	r, e := c.generated.DeleteSessionByIdentifierWithResponse(
		x,
		identifier,
		params,
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"delete session by identifier",
			r.HTTPResponse,
			r.Body,
		)
	}

	return r.JSON200, nil
}
