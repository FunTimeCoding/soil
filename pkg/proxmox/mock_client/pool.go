package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) Pool(name string) (*proxmox.Pool, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	return &proxmox.Pool{PoolID: name}, nil
}
