package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) DeleteMachineSnapshot(
	identifier int64,
	name string,
	node *string,
) (string, int) {
	result, e := c.client.DeleteMachineSnapshotWithResponse(
		c.context,
		identifier,
		name,
		&client.DeleteMachineSnapshotParams{
			Instance: &c.instance,
			Node:     node,
		},
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
