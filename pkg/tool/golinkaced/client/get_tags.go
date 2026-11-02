package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetTags(
	x context.Context,
	p *client.GetTagsParams,
) (*client.TagPage, error) {
	r, e := c.client.GetTagsWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get tags", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
