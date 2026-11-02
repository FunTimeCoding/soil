package mock_client

import (
	"context"
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) CloneMachine(
	_ context.Context,
	identifier int64,
	_ *client.CloneMachineParams,
	b client.CloneMachineJSONRequestBody,
) (*client.CreateMachineResult, error) {
	i, e := c.machine(identifier)

	if e != nil {
		return nil, e
	}

	target := int64(100 + len(c.MachineList))

	if b.NewIdentifier != nil {
		target = *b.NewIdentifier
	}

	status := proxmoxConstant.StoppedStatus
	c.MachineList = append(
		c.MachineList,
		client.Machine{
			Identifier: target,
			Name:       b.Name,
			Node:       c.MachineList[i].Node,
			Status:     &status,
		},
	)

	return &client.CreateMachineResult{Identifier: target, Status: &status}, nil
}
