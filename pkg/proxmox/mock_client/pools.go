package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) Pools() ([]*proxmox.Pool, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	return []*proxmox.Pool{}, nil
}
