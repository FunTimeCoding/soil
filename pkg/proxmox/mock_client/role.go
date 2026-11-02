package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) Role(_ string) (*proxmox.Permission, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	return &proxmox.Permission{}, nil
}
