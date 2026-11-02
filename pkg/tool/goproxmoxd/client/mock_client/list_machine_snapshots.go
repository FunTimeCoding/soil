package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListMachineSnapshots(
	_ context.Context,
	identifier int64,
	_ *client.ListMachineSnapshotsParams,
) ([]client.Snapshot, error) {
	if _, e := c.machine(identifier); e != nil {
		return nil, e
	}

	return c.MachineSnapshots[identifier], nil
}
