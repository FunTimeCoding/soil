package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) DeleteContainerSnapshot(
	v *proxmox.Container,
	name string,
) (*proxmox.Task, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	identifier := int(v.VMID)
	var remaining []*proxmox.ContainerSnapshot

	for _, s := range c.containerSnapshots[identifier] {
		if s.Name != name {
			remaining = append(remaining, s)
		}
	}

	c.containerSnapshots[identifier] = remaining

	return &proxmox.Task{UPID: "mock:ct-delete-snapshot"}, nil
}
