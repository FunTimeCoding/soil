package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) CreateContainerSnapshot(
	v *proxmox.Container,
	name string,
) (*proxmox.Task, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	identifier := int(v.VMID)
	c.containerSnapshots[identifier] = append(
		c.containerSnapshots[identifier],
		&proxmox.ContainerSnapshot{Name: name},
	)

	return &proxmox.Task{UPID: "mock:ct-snapshot"}, nil
}
