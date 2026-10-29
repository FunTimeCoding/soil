package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetTags(x context.Context) ([]*client.SourceTypeTag, error) {
	r, e := c.client.GetTagsWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list tags", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
