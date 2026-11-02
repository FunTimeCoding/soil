package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) Group(name string) (*proxmox.Group, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	return &proxmox.Group{GroupID: name}, nil
}
