package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) DeleteMachineSnapshot(
	v *proxmox.VirtualMachine,
	name string,
) (*proxmox.Task, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	identifier := int(v.VMID)
	var remaining []*proxmox.VirtualMachineSnapshot

	for _, s := range c.machineSnapshots[identifier] {
		if s.Name != name {
			remaining = append(remaining, s)
		}
	}

	c.machineSnapshots[identifier] = remaining

	return &proxmox.Task{UPID: "mock:delete-snapshot"}, nil
}
