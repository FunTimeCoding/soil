package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) UpdateMachine(
	_ context.Context,
	identifier int64,
	_ *client.UpdateMachineParams,
	b client.UpdateMachineJSONRequestBody,
) (*client.StatusResult, error) {
	i, e := c.machine(identifier)

	if e != nil {
		return nil, e
	}

	if b.Name != nil {
		c.MachineList[i].Name = *b.Name
	}

	if b.Tags != nil {
		c.MachineList[i].Tags = b.Tags
	}

	return &client.StatusResult{Status: "updated"}, nil
}
