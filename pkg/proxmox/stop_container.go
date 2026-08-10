package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) StopContainer(v *proxmox.Container) (*proxmox.Task, error) {
	return v.Stop(c.context)
}
