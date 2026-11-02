package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) DeleteMachine(
	v *proxmox.VirtualMachine,
	_ *proxmox.VirtualMachineDeleteOptions,
) (*proxmox.Task, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	nodeMachines, okay := c.machines[v.Node]

	if !okay {
		return nil, not_found.New(constant.MachineKind, v.VMID)
	}

	delete(nodeMachines, int(v.VMID))

	return &proxmox.Task{UPID: "mock:delete"}, nil
}
