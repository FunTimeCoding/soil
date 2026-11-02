package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) Domains() (proxmox.Domains, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	return proxmox.Domains{}, nil
}
