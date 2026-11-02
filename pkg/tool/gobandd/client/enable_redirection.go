package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) EnableRedirection(
	x context.Context,
	p client.RedirectionRequest,
) (*client.RedirectionResult, error) {
	r, e := c.client.EnableRedirectionWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("redirection", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
