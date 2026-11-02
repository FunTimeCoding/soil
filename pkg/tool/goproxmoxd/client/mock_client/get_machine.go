package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) GetMachine(
	_ context.Context,
	identifier int64,
	_ *client.GetMachineParams,
) (*client.MachineDetail, error) {
	i, e := c.machine(identifier)

	if e != nil {
		return nil, e
	}

	m := c.MachineList[i]

	return &client.MachineDetail{
		Identifier: m.Identifier,
		Name:       m.Name,
		Node:       m.Node,
		Status:     m.Status,
		Tags:       m.Tags,
	}, nil
}
