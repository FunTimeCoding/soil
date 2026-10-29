package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListClusters(x context.Context) ([]*client.Cluster, error) {
	r, e := c.client.ListClustersWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list clusters", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
