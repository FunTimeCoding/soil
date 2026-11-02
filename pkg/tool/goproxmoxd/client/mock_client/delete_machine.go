package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"slices"
)

func (c *Client) DeleteMachine(
	_ context.Context,
	identifier int64,
	_ *client.DeleteMachineParams,
) (*client.TaskResult, error) {
	i, e := c.machine(identifier)

	if e != nil {
		return nil, e
	}

	c.MachineList = slices.Delete(c.MachineList, i, i+1)
	delete(c.MachineSnapshots, identifier)

	return c.task(), nil
}
