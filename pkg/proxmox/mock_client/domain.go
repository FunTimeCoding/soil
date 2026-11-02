package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) Domain(name string) (*proxmox.Domain, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	return &proxmox.Domain{Realm: name}, nil
}
