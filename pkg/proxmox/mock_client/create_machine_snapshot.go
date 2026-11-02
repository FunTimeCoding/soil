package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) CreateMachineSnapshot(
	v *proxmox.VirtualMachine,
	name string,
) (*proxmox.Task, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	identifier := int(v.VMID)
	c.machineSnapshots[identifier] = append(
		c.machineSnapshots[identifier],
		&proxmox.VirtualMachineSnapshot{Name: name},
	)

	return &proxmox.Task{UPID: "mock:snapshot"}, nil
}
