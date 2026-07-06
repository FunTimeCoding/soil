package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) DeleteContainerSnapshot(
	identifier int64,
	name string,
	node *string,
) string {
	result, e := c.client.DeleteContainerSnapshotWithResponse(
		c.context,
		identifier,
		name,
		&client.DeleteContainerSnapshotParams{
			Instance: &c.instance,
			Node:     node,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
