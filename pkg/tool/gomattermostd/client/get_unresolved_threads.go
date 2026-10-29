package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetUnresolvedThreads(
	x context.Context,
	name string,
	p *client.GetUnresolvedThreadsParams,
) ([]client.Thread, error) {
	r, e := c.client.GetUnresolvedThreadsWithResponse(x, name, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("unresolved threads", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
