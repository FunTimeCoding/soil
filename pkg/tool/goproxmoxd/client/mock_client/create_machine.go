package mock_client

import (
	"context"
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) CreateMachine(
	_ context.Context,
	_ *client.CreateMachineParams,
	b client.CreateMachineJSONRequestBody,
) (*client.CreateMachineResult, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	identifier := int64(100 + len(c.MachineList))

	if b.Identifier != nil {
		identifier = *b.Identifier
	}

	status := proxmoxConstant.StoppedStatus
	c.MachineList = append(
		c.MachineList,
		client.Machine{
			Identifier: identifier,
			Name:       b.Name,
			Node:       &b.Node,
			Status:     &status,
			Tags:       b.Tags,
		},
	)

	return &client.CreateMachineResult{
		Identifier: identifier,
		Status:     &status,
	}, nil
}
