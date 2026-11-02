package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) NodeVersion(_ *proxmox.Node) (*proxmox.Version, error) {
	return c.version, nil
}
