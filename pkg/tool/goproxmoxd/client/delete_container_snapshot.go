package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) DeleteContainerSnapshot(
	identifier int64,
	name string,
	node *string,
) (string, int) {
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

	return string(result.Body), result.StatusCode()
}
