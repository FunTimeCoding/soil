package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) CloneMachine(
	vm *proxmox.VirtualMachine,
	options *proxmox.VirtualMachineCloneOptions,
) (int, *proxmox.Task, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	newIdentifier := options.NewID

	if newIdentifier == 0 {
		newIdentifier = c.nextIdentifier
		c.nextIdentifier++
	}

	clone := &proxmox.VirtualMachine{
		VMID: proxmox.StringOrUint64(newIdentifier),
		Name: options.Name,
		Node: vm.Node,
	}
	c.machines[vm.Node][newIdentifier] = clone

	return newIdentifier, &proxmox.Task{UPID: "mock:clone"}, nil
}
