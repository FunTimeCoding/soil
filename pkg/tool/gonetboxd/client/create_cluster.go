package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateCluster(
	x context.Context,
	name string,
	clusterType string,
	site string,
) (client.Cluster, error) {
	r, e := c.client.CreateClusterWithResponse(
		x,
		client.CreateClusterRequest{Name: name, Type: clusterType, Site: site},
	)

	if e != nil {
		return client.Cluster{}, e
	}

	if r.JSON201 == nil {
		return client.Cluster{}, refusal.New(
			"create cluster",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
