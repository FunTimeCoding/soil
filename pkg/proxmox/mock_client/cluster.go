package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) Cluster() (*proxmox.Cluster, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	return &proxmox.Cluster{}, nil
}
