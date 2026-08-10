package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) StartContainer(v *proxmox.Container) (*proxmox.Task, error) {
	return v.Start(c.context)
}
