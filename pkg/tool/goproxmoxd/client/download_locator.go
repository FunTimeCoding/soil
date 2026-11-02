package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) DownloadLocator(
	x context.Context,
	node string,
	storage string,
	p *client.DownloadLocatorParams,
	b client.DownloadLocatorJSONRequestBody,
) (*client.StatusResult, error) {
	r, e := c.client.DownloadLocatorWithResponse(x, node, storage, p, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("download locator", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
