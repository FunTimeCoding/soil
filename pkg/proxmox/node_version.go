package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) NodeVersion(n *proxmox.Node) (*proxmox.Version, error) {
	return n.Version(c.context)
}
