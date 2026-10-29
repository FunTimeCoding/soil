package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateClusterType(
	x context.Context,
	name string,
) (client.ClusterType, error) {
	r, e := c.client.CreateClusterTypeWithResponse(
		x,
		client.CreateNameRequest{Name: name},
	)

	if e != nil {
		return client.ClusterType{}, e
	}

	if r.JSON201 == nil {
		return client.ClusterType{}, refusal.New(
			"create cluster type",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
