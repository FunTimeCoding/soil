package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateCluster(
	name string,
	clusterType string,
	site string,
) string {
	result, e := c.client.CreateCluster(
		c.context,
		client.CreateClusterRequest{
			Name: name,
			Type: clusterType,
			Site: site,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
