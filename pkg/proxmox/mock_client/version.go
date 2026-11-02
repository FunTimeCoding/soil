package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) Version() (*proxmox.Version, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	return c.version, nil
}
