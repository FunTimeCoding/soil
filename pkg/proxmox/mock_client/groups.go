package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) Groups() (proxmox.Groups, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	return proxmox.Groups{}, nil
}
