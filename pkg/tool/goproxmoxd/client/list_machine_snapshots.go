package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListMachineSnapshots(
	identifier int64,
	node *string,
) (string, int) {
	result, e := c.client.ListMachineSnapshotsWithResponse(
		c.context,
		identifier,
		&client.ListMachineSnapshotsParams{Instance: &c.instance, Node: node},
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
