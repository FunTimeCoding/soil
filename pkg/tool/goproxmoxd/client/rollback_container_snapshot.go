package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) RollbackContainerSnapshot(
	identifier int64,
	name string,
	node *string,
) (string, int) {
	result, e := c.client.RollbackContainerSnapshotWithResponse(
		c.context,
		identifier,
		name,
		&client.RollbackContainerSnapshotParams{
			Instance: &c.instance,
			Node:     node,
		},
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
