package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) CreateMachineSnapshot(
	_ context.Context,
	identifier int64,
	_ *client.CreateMachineSnapshotParams,
	b client.CreateMachineSnapshotJSONRequestBody,
) (*client.TaskResult, error) {
	if _, e := c.machine(identifier); e != nil {
		return nil, e
	}

	c.MachineSnapshots[identifier] = append(
		c.MachineSnapshots[identifier],
		client.Snapshot{Name: b.Name},
	)

	return c.task(), nil
}
