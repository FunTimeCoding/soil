package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListLibraries(x context.Context) ([]*client.Library, error) {
	r, e := c.client.ListLibrariesWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list libraries", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
