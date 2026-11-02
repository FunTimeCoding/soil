package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) MachineSnapshots(v *proxmox.VirtualMachine) ([]*proxmox.VirtualMachineSnapshot, error) {
	return c.machineSnapshots[int(v.VMID)], nil
}
