package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) ContainerSnapshots(v *proxmox.Container) ([]*proxmox.ContainerSnapshot, error) {
	return c.containerSnapshots[int(v.VMID)], nil
}
