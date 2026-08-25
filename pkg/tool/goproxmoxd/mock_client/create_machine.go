package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) CreateMachine(
	n *proxmox.Node,
	identifier int,
	options ...proxmox.VirtualMachineOption,
) (*proxmox.Task, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	nodeMachines, okay := c.machines[n.Name]

	if !okay {
		return nil, fmt.Errorf("node %q not found", n.Name)
	}

	vm := &proxmox.VirtualMachine{
		VMID: proxmox.StringOrUint64(identifier),
		Name: optionValue(options, constant.NameOption),
		Node: n.Name,
	}
	nodeMachines[identifier] = vm

	return &proxmox.Task{UPID: "mock:create"}, nil
}
