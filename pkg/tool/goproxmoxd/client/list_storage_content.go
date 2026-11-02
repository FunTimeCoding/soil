package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListStorageContent(
	x context.Context,
	node string,
	storage string,
	p *client.ListStorageContentParams,
) ([]client.StorageContentItem, error) {
	r, e := c.client.ListStorageContentWithResponse(x, node, storage, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list storage content", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
