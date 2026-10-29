package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostIndex(
	x context.Context,
	b client.PostIndexJSONRequestBody,
) ([]*client.IndexResult, error) {
	r, e := c.client.PostIndexWithResponse(x, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(constant.Index, r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
