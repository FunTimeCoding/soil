package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"slices"
)

func (c *Client) DeleteMachineSnapshot(
	_ context.Context,
	identifier int64,
	name string,
	_ *client.DeleteMachineSnapshotParams,
) (*client.TaskResult, error) {
	if _, e := c.machine(identifier); e != nil {
		return nil, e
	}

	i, e := snapshot(c.MachineSnapshots[identifier], name)

	if e != nil {
		return nil, e
	}

	c.MachineSnapshots[identifier] = slices.Delete(
		c.MachineSnapshots[identifier],
		i,
		i+1,
	)

	return c.task(), nil
}
