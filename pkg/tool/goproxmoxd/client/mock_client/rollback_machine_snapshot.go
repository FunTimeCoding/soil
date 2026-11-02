package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) RollbackMachineSnapshot(
	_ context.Context,
	identifier int64,
	name string,
	_ *client.RollbackMachineSnapshotParams,
) (*client.TaskResult, error) {
	if _, e := c.machine(identifier); e != nil {
		return nil, e
	}

	if _, e := snapshot(c.MachineSnapshots[identifier], name); e != nil {
		return nil, e
	}

	return c.task(), nil
}
