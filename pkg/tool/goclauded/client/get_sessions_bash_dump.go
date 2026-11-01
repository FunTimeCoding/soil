package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetSessionsBashDump(
	x context.Context,
) (*client.BashDumpResponse, error) {
	r, e := c.generated.GetSessionsBashDumpWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"get sessions bash dump",
			r.HTTPResponse,
			r.Body,
		)
	}

	return r.JSON200, nil
}
