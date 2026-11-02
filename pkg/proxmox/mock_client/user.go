package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) User(name string) (*proxmox.User, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	return &proxmox.User{UserID: name}, nil
}
