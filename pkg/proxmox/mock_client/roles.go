package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) Roles() ([]*proxmox.Role, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	return []*proxmox.Role{}, nil
}
